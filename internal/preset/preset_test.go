package preset

import (
	"fmt"
	"testing"

	"github.com/Ooooze/batctl/internal/backend"
)

type mockBackend struct {
	name string
	caps backend.Capabilities
}

func (m *mockBackend) Name() string                { return m.name }
func (m *mockBackend) Detect() bool                { return false }
func (m *mockBackend) Capabilities() backend.Capabilities { return m.caps }
func (m *mockBackend) GetThresholds(bat string) (int, int, error) {
	return 0, 0, fmt.Errorf("mock")
}
func (m *mockBackend) SetThresholds(bat string, start, stop int) error {
	return fmt.Errorf("mock")
}
func (m *mockBackend) GetChargeBehaviour(bat string) (string, []string, error) {
	return "", nil, fmt.Errorf("mock")
}
func (m *mockBackend) SetChargeBehaviour(bat string, mode string) error {
	return fmt.Errorf("mock")
}

func (m *mockBackend) ValidateThresholds(start, stop int) error {
	caps := m.caps
	if caps.StartThreshold {
		if start < caps.StartRange[0] || start > caps.StartRange[1] {
			return fmt.Errorf("start out of range")
		}
	}
	if caps.StopThreshold {
		if len(caps.DiscreteStopVals) > 0 {
			valid := false
			for _, v := range caps.DiscreteStopVals {
				if stop == v {
					valid = true
					break
				}
			}
			if !valid {
				return fmt.Errorf("stop not in discrete values")
			}
		} else {
			if stop < caps.StopRange[0] || stop > caps.StopRange[1] {
				return fmt.Errorf("stop out of range")
			}
		}
	}
	if caps.StartThreshold && start >= stop {
		return fmt.Errorf("start >= stop")
	}
	return nil
}

func TestFindByID(t *testing.T) {
	tests := []struct {
		id    string
		found bool
	}{
		{"max-lifespan", true},
		{"balanced", true},
		{"full-charge", true},
		{"plugged-in", true},
		{"nonexistent", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			p, ok := FindByID(tt.id)
			if ok != tt.found {
				t.Fatalf("FindByID(%q) found=%v, want %v", tt.id, ok, tt.found)
			}
			if ok && p.ID != tt.id {
				t.Fatalf("FindByID(%q) returned preset with ID=%q", tt.id, p.ID)
			}
		})
	}
}

func TestNearestDiscrete(t *testing.T) {
	tests := []struct {
		name   string
		target int
		vals   []int
		want   int
	}{
		{"exact match", 80, []int{50, 80, 100}, 80},
		{"closest to 50", 60, []int{50, 80, 100}, 50},
		{"closest to 80", 75, []int{50, 80, 100}, 80},
		{"closest to 100", 95, []int{50, 80, 100}, 100},
		{"boundary low", 0, []int{50, 80, 100}, 50},
		{"boundary high", 200, []int{50, 80, 100}, 100},
		{"single value", 42, []int{80}, 80},
		{"equidistant keeps first found", 65, []int{50, 80, 100}, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nearestDiscrete(tt.target, tt.vals)
			if got != tt.want {
				t.Fatalf("nearestDiscrete(%d, %v) = %d, want %d", tt.target, tt.vals, got, tt.want)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		name     string
		v, min, max, want int
	}{
		{"within range", 50, 0, 100, 50},
		{"at min", 0, 0, 100, 0},
		{"at max", 100, 0, 100, 100},
		{"below min", -5, 0, 100, 0},
		{"above max", 150, 0, 100, 100},
		{"min equals max", 50, 80, 80, 80},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clamp(tt.v, tt.min, tt.max)
			if got != tt.want {
				t.Fatalf("clamp(%d, %d, %d) = %d, want %d", tt.v, tt.min, tt.max, got, tt.want)
			}
		})
	}
}

func TestAdaptToBackend(t *testing.T) {
	thinkpadCaps := backend.Capabilities{
		StartThreshold: true, StopThreshold: true,
		StartRange: [2]int{0, 99}, StopRange: [2]int{1, 100},
	}
	asusCaps := backend.Capabilities{
		StartThreshold: false, StopThreshold: true,
		StopRange: [2]int{0, 100},
	}
	sonyCaps := backend.Capabilities{
		StartThreshold: false, StopThreshold: true,
		DiscreteStopVals: []int{50, 80, 100},
	}
	dellCaps := backend.Capabilities{
		StartThreshold: true, StopThreshold: true,
		StartRange: [2]int{50, 95}, StopRange: [2]int{55, 100},
	}

	tests := []struct {
		name      string
		preset    Preset
		backend   backend.Backend
		wantStart int
		wantStop  int
		wantErr   bool
	}{
		{
			"ThinkPad balanced",
			Preset{ID: "balanced", Start: 40, Stop: 80},
			&mockBackend{name: "ThinkPad", caps: thinkpadCaps},
			40, 80, false,
		},
		{
			"ThinkPad max-lifespan",
			Preset{ID: "max-lifespan", Start: 20, Stop: 80},
			&mockBackend{name: "ThinkPad", caps: thinkpadCaps},
			20, 80, false,
		},
		{
			"ASUS max-lifespan (no start threshold)",
			Preset{ID: "max-lifespan", Start: 20, Stop: 80},
			&mockBackend{name: "ASUS", caps: asusCaps},
			0, 80, false,
		},
		{
			"Sony balanced (discrete stop)",
			Preset{ID: "balanced", Start: 40, Stop: 80},
			&mockBackend{name: "Sony", caps: sonyCaps},
			0, 80, false,
		},
		{
			"Dell max-lifespan (clamped start)",
			Preset{ID: "max-lifespan", Start: 20, Stop: 80},
			&mockBackend{name: "Dell", caps: dellCaps},
			50, 80, false,
		},
		{
			"ThinkPad full-charge",
			Preset{ID: "full-charge", Start: 0, Stop: 100},
			&mockBackend{name: "ThinkPad", caps: thinkpadCaps},
			0, 100, false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, stop, err := AdaptToBackend(tt.preset, tt.backend)
			if (err != nil) != tt.wantErr {
				t.Fatalf("AdaptToBackend() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if start != tt.wantStart || stop != tt.wantStop {
				t.Fatalf("AdaptToBackend() = %d/%d, want %d/%d", start, stop, tt.wantStart, tt.wantStop)
			}
		})
	}
}
