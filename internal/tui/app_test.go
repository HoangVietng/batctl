package tui

import "testing"

func TestNextDiscrete(t *testing.T) {
	vals := []int{60, 80, 100}

	tests := []struct {
		name      string
		current   int
		direction int
		vals      []int
		want      int
	}{
		{"step up from 80", 80, 1, vals, 100},
		{"step down from 80", 80, -1, vals, 60},
		{"step up from 100 (clamped)", 100, 1, vals, 100},
		{"step down from 60 (clamped)", 60, -1, vals, 60},
		{"step up from 60", 60, 1, vals, 80},
		{"step down from 100", 100, -1, vals, 80},
		{"single value step up", 80, 1, []int{80}, 80},
		{"single value step down", 80, -1, []int{80}, 80},
		{"step up by 5 from 60", 60, 5, vals, 100},
		{"step down by -5 from 100", 100, -5, vals, 60},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nextDiscrete(tt.current, tt.direction, tt.vals)
			if got != tt.want {
				t.Fatalf("nextDiscrete(%d, %d, %v) = %d, want %d",
					tt.current, tt.direction, tt.vals, got, tt.want)
			}
		})
	}
}
