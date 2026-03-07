package backend

import (
	"fmt"
	"strings"

	"github.com/Ooooze/batctl/internal/battery"
)

const acerHealthModePath = "/sys/bus/wmi/drivers/acer-wmi-battery/health_mode"

type AcerBackend struct{}

func (b *AcerBackend) Name() string {
	return "Acer"
}

func (b *AcerBackend) Detect() bool {
	vendor := DetectVendor()
	if !strings.Contains(vendor, "Acer") {
		return false
	}
	return battery.SysfsExists(acerHealthModePath)
}

func (b *AcerBackend) Capabilities() Capabilities {
	return Capabilities{
		StartThreshold:    false,
		StopThreshold:     true,
		ChargeBehaviour:   false,
		StartRange:        [2]int{0, 0},
		StopRange:         [2]int{80, 100},
		DiscreteStopVals:  []int{80, 100},
		StartAutoComputed: false,
	}
}

func (b *AcerBackend) GetThresholds(bat string) (start, stop int, err error) {
	val, err := battery.SysfsReadInt(acerHealthModePath)
	if err != nil {
		return 0, 0, err
	}
	if val == 1 {
		return 0, 80, nil
	}
	return 0, 100, nil
}

func (b *AcerBackend) SetThresholds(bat string, start, stop int) error {
	if err := b.ValidateThresholds(start, stop); err != nil {
		return err
	}
	if stop <= 80 {
		return battery.SysfsWriteInt(acerHealthModePath, 1)
	}
	return battery.SysfsWriteInt(acerHealthModePath, 0)
}

func (b *AcerBackend) GetChargeBehaviour(bat string) (current string, available []string, err error) {
	return "", nil, fmt.Errorf("not supported")
}

func (b *AcerBackend) SetChargeBehaviour(bat string, mode string) error {
	return fmt.Errorf("not supported")
}

func (b *AcerBackend) ValidateThresholds(start, stop int) error {
	if stop != 80 && stop != 100 {
		return fmt.Errorf("stop threshold must be 80 or 100, got %d", stop)
	}
	return nil
}
