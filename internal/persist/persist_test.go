package persist

import (
	"os"
	"path/filepath"
	"testing"
)

func withTempConfig(t *testing.T) func() {
	t.Helper()
	dir := t.TempDir()
	orig := ConfigPath
	ConfigPath = filepath.Join(dir, "batctl.conf")
	return func() { ConfigPath = orig }
}

func TestSaveAndLoadConfig(t *testing.T) {
	restore := withTempConfig(t)
	defer restore()

	cfg := Config{Battery: "BAT0", Start: 40, Stop: 80}
	if err := SaveConfig(cfg); err != nil {
		t.Fatalf("SaveConfig: %v", err)
	}

	got, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig: %v", err)
	}

	if got.Battery != cfg.Battery {
		t.Fatalf("Battery = %q, want %q", got.Battery, cfg.Battery)
	}
	if got.Start != cfg.Start {
		t.Fatalf("Start = %d, want %d", got.Start, cfg.Start)
	}
	if got.Stop != cfg.Stop {
		t.Fatalf("Stop = %d, want %d", got.Stop, cfg.Stop)
	}
}

func TestSaveAndLoadConfigMultiBattery(t *testing.T) {
	restore := withTempConfig(t)
	defer restore()

	cfg := Config{Battery: "all", Start: 20, Stop: 80}
	if err := SaveConfig(cfg); err != nil {
		t.Fatalf("SaveConfig: %v", err)
	}

	got, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig: %v", err)
	}

	if got.Battery != "all" {
		t.Fatalf("Battery = %q, want %q", got.Battery, "all")
	}
	if got.Start != 20 || got.Stop != 80 {
		t.Fatalf("thresholds = %d/%d, want 20/80", got.Start, got.Stop)
	}
}

func TestLoadConfigMissing(t *testing.T) {
	restore := withTempConfig(t)
	defer restore()

	_, err := LoadConfig()
	if err == nil {
		t.Fatal("expected error for missing config file")
	}
}

func TestLoadConfigComments(t *testing.T) {
	restore := withTempConfig(t)
	defer restore()

	content := `# This is a comment
# Another comment

BATTERY=BAT1
START_THRESHOLD=50

# inline comment
STOP_THRESHOLD=90
`
	if err := os.WriteFile(ConfigPath, []byte(content), 0644); err != nil {
		t.Fatalf("writing test config: %v", err)
	}

	got, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig: %v", err)
	}

	if got.Battery != "BAT1" {
		t.Fatalf("Battery = %q, want %q", got.Battery, "BAT1")
	}
	if got.Start != 50 {
		t.Fatalf("Start = %d, want 50", got.Start)
	}
	if got.Stop != 90 {
		t.Fatalf("Stop = %d, want 90", got.Stop)
	}
}

func TestLoadConfigDefaults(t *testing.T) {
	restore := withTempConfig(t)
	defer restore()

	if err := os.WriteFile(ConfigPath, []byte("STOP_THRESHOLD=80\n"), 0644); err != nil {
		t.Fatalf("writing test config: %v", err)
	}

	got, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig: %v", err)
	}

	if got.Battery != "BAT0" {
		t.Fatalf("Battery default = %q, want %q", got.Battery, "BAT0")
	}
	if got.Start != 0 {
		t.Fatalf("Start default = %d, want 0", got.Start)
	}
	if got.Stop != 80 {
		t.Fatalf("Stop = %d, want 80", got.Stop)
	}
}
