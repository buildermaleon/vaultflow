package config

import (
	"os"
	"testing"
)

func TestDefault(t *testing.T) {
	cfg := Default()
	if cfg == nil {
		t.Fatal("Default() returned nil")
	}
	if cfg.MasterKey != "" {
		t.Errorf("Default MasterKey = %v, want empty", cfg.MasterKey)
	}
	if cfg.VaultFile == "" {
		t.Error("Default VaultFile should not be empty")
	}
}

func TestDefaultWithEnv(t *testing.T) {
	os.Setenv("VAULT_MASTER_KEY", "test-key")
	defer os.Unsetenv("VAULT_MASTER_KEY")
	
	cfg := Default()
	if cfg.MasterKey != "test-key" {
		t.Errorf("MasterKey = %v, want test-key", cfg.MasterKey)
	}
}
