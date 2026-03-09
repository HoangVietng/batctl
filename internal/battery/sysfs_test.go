package battery

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBatPath(t *testing.T) {
	tests := []struct {
		bat, attr, want string
	}{
		{"BAT0", "status", "/sys/class/power_supply/BAT0/status"},
		{"BAT1", "charge_control_end_threshold", "/sys/class/power_supply/BAT1/charge_control_end_threshold"},
		{"BAT0", "capacity", "/sys/class/power_supply/BAT0/capacity"},
	}

	for _, tt := range tests {
		t.Run(tt.bat+"/"+tt.attr, func(t *testing.T) {
			got := BatPath(tt.bat, tt.attr)
			if got != tt.want {
				t.Fatalf("BatPath(%q, %q) = %q, want %q", tt.bat, tt.attr, got, tt.want)
			}
		})
	}
}

func TestSysfsReadString(t *testing.T) {
	dir := t.TempDir()

	t.Run("reads and trims", func(t *testing.T) {
		path := filepath.Join(dir, "test_attr")
		os.WriteFile(path, []byte("hello\n"), 0644)
		got, err := SysfsReadString(path)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "hello" {
			t.Fatalf("got %q, want %q", got, "hello")
		}
	})

	t.Run("trims spaces and newlines", func(t *testing.T) {
		path := filepath.Join(dir, "test_spaces")
		os.WriteFile(path, []byte("  Charging  \n"), 0644)
		got, err := SysfsReadString(path)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "Charging" {
			t.Fatalf("got %q, want %q", got, "Charging")
		}
	})

	t.Run("returns error for missing file", func(t *testing.T) {
		_, err := SysfsReadString(filepath.Join(dir, "nonexistent"))
		if err == nil {
			t.Fatal("expected error for missing file")
		}
	})
}

func TestSysfsReadInt(t *testing.T) {
	dir := t.TempDir()

	t.Run("reads integer", func(t *testing.T) {
		path := filepath.Join(dir, "int_val")
		os.WriteFile(path, []byte("42\n"), 0644)
		got, err := SysfsReadInt(path)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 42 {
			t.Fatalf("got %d, want 42", got)
		}
	})

	t.Run("reads zero", func(t *testing.T) {
		path := filepath.Join(dir, "zero_val")
		os.WriteFile(path, []byte("0\n"), 0644)
		got, err := SysfsReadInt(path)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 0 {
			t.Fatalf("got %d, want 0", got)
		}
	})

	t.Run("returns error for non-integer", func(t *testing.T) {
		path := filepath.Join(dir, "bad_val")
		os.WriteFile(path, []byte("abc\n"), 0644)
		_, err := SysfsReadInt(path)
		if err == nil {
			t.Fatal("expected error for non-integer content")
		}
	})

	t.Run("returns error for missing file", func(t *testing.T) {
		_, err := SysfsReadInt(filepath.Join(dir, "nonexistent"))
		if err == nil {
			t.Fatal("expected error for missing file")
		}
	})
}

func TestSysfsWriteInt(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "write_int")

	if err := SysfsWriteInt(path, 80); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read back: %v", err)
	}
	if string(data) != "80" {
		t.Fatalf("got %q, want %q", string(data), "80")
	}
}

func TestSysfsWriteString(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "write_str")

	if err := SysfsWriteString(path, "auto"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read back: %v", err)
	}
	if string(data) != "auto" {
		t.Fatalf("got %q, want %q", string(data), "auto")
	}
}

func TestSysfsExists(t *testing.T) {
	dir := t.TempDir()

	t.Run("existing file", func(t *testing.T) {
		path := filepath.Join(dir, "exists")
		os.WriteFile(path, []byte(""), 0644)
		if !SysfsExists(path) {
			t.Fatal("expected true for existing file")
		}
	})

	t.Run("missing file", func(t *testing.T) {
		if SysfsExists(filepath.Join(dir, "nope")) {
			t.Fatal("expected false for missing file")
		}
	})
}
