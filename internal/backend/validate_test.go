package backend

import "testing"

func TestValidateThresholds(t *testing.T) {
	tests := []struct {
		name    string
		backend Backend
		start   int
		stop    int
		wantErr bool
	}{
		// ThinkPad: start 0-99, stop 1-100, start < stop
		{"ThinkPad valid", &ThinkPadBackend{}, 40, 80, false},
		{"ThinkPad min values", &ThinkPadBackend{}, 0, 1, false},
		{"ThinkPad max values", &ThinkPadBackend{}, 99, 100, false},
		{"ThinkPad start too high", &ThinkPadBackend{}, 100, 100, true},
		{"ThinkPad stop too low", &ThinkPadBackend{}, 0, 0, true},
		{"ThinkPad start >= stop", &ThinkPadBackend{}, 80, 80, true},
		{"ThinkPad negative start", &ThinkPadBackend{}, -1, 80, true},
		{"ThinkPad stop > 100", &ThinkPadBackend{}, 40, 101, true},

		// ASUS: stop 0-100 only
		{"ASUS valid", &ASUSBackend{}, 0, 80, false},
		{"ASUS stop 0", &ASUSBackend{}, 0, 0, false},
		{"ASUS stop 100", &ASUSBackend{}, 0, 100, false},
		{"ASUS stop negative", &ASUSBackend{}, 0, -1, true},
		{"ASUS stop > 100", &ASUSBackend{}, 0, 101, true},

		// Dell: start 50-95, stop 55-100, stop - start == 5
		{"Dell valid", &DellBackend{}, 75, 80, false},
		{"Dell min values", &DellBackend{}, 50, 55, false},
		{"Dell max values", &DellBackend{}, 95, 100, false},
		{"Dell start too low", &DellBackend{}, 49, 54, true},
		{"Dell stop too low", &DellBackend{}, 50, 54, true},
		{"Dell start >= stop", &DellBackend{}, 80, 80, true},
		{"Dell gap not 5", &DellBackend{}, 50, 60, true},

		// Sony: stop in {50, 80, 100}
		{"Sony stop 50", &SonyBackend{}, 0, 50, false},
		{"Sony stop 80", &SonyBackend{}, 0, 80, false},
		{"Sony stop 100", &SonyBackend{}, 0, 100, false},
		{"Sony stop 60 invalid", &SonyBackend{}, 0, 60, true},
		{"Sony stop 0 invalid", &SonyBackend{}, 0, 0, true},

		// Acer: stop in {80, 100}
		{"Acer stop 80", &AcerBackend{}, 0, 80, false},
		{"Acer stop 100", &AcerBackend{}, 0, 100, false},
		{"Acer stop 60 invalid", &AcerBackend{}, 0, 60, true},

		// MSI: stop 10-100
		{"MSI valid", &MSIBackend{}, 0, 80, false},
		{"MSI stop 10", &MSIBackend{}, 0, 10, false},
		{"MSI stop 100", &MSIBackend{}, 0, 100, false},
		{"MSI stop 9 too low", &MSIBackend{}, 0, 9, true},
		{"MSI stop > 100", &MSIBackend{}, 0, 101, true},

		// Tuxedo: start in {40,50,60,70,80,95}, stop in {60,70,80,90,100}
		{"Tuxedo valid", &TuxedoBackend{}, 40, 60, false},
		{"Tuxedo valid 2", &TuxedoBackend{}, 80, 100, false},
		{"Tuxedo invalid start", &TuxedoBackend{}, 45, 80, true},
		{"Tuxedo invalid stop", &TuxedoBackend{}, 40, 65, true},
		{"Tuxedo start >= stop", &TuxedoBackend{}, 80, 80, true},

		// Generic: start 0-99, stop 1-100, start < stop
		{"Generic valid", &GenericBackend{caps: &Capabilities{
			StartThreshold: true, StopThreshold: true,
			StartRange: [2]int{0, 99}, StopRange: [2]int{1, 100},
		}}, 40, 80, false},
		{"Generic start >= stop", &GenericBackend{caps: &Capabilities{
			StartThreshold: true, StopThreshold: true,
			StartRange: [2]int{0, 99}, StopRange: [2]int{1, 100},
		}}, 80, 80, true},

		// System76: start 0-99, stop 1-100, start < stop
		{"System76 valid", &System76Backend{}, 40, 80, false},
		{"System76 start >= stop", &System76Backend{}, 80, 80, true},
		{"System76 start negative", &System76Backend{}, -1, 80, true},

		// Huawei: start 0-99, stop 1-100, start < stop
		{"Huawei valid", &HuaweiBackend{}, 40, 80, false},
		{"Huawei start >= stop", &HuaweiBackend{}, 80, 80, true},

		// Surface: stop 1-100 (start depends on sysfs availability)
		{"Surface valid (stop only)", &SurfaceBackend{}, 0, 80, false},
		{"Surface stop too low", &SurfaceBackend{}, 0, 0, true},
		{"Surface stop > 100", &SurfaceBackend{}, 0, 101, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.backend.ValidateThresholds(tt.start, tt.stop)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateThresholds(%d, %d) error = %v, wantErr %v",
					tt.start, tt.stop, err, tt.wantErr)
			}
		})
	}
}
