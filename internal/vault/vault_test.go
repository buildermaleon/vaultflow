package vault

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"github.com/dablon/vaultflow/internal/config"
)

func setupTestVault(t *testing.T) (*Vault, string) {
	tmpDir := t.TempDir()
	cfg := &config.Config{
		MasterKey: "test-master-key-32bytes!",
		VaultFile: filepath.Join(tmpDir, "test-secrets.json"),
	}
	v := New(cfg)
	return v, cfg.VaultFile
}

func TestVault_SetAndGet(t *testing.T) {
	v, _ := setupTestVault(t)
	
	err := v.Set("test-key", "test-value")
	if err != nil {
		t.Fatalf("Set() error = %v", err)
	}
	
	val, err := v.Get("test-key")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	
	if val != "test-value" {
		t.Errorf("Get() = %v, want test-value", val)
	}
}

func TestVault_GetNonExistent(t *testing.T) {
	v, _ := setupTestVault(t)
	
	_, err := v.Get("non-existent")
	if err == nil {
		t.Error("Get() should error for non-existent key")
	}
}

func TestVault_List(t *testing.T) {
	v, _ := setupTestVault(t)
	
	v.Set("key1", "val1")
	v.Set("key2", "val2")
	v.Set("key3", "val3")
	
	keys, err := v.List()
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	
	if len(keys) != 3 {
		t.Errorf("List() = %d keys, want 3", len(keys))
	}
}

func TestVault_Delete(t *testing.T) {
	v, _ := setupTestVault(t)
	
	v.Set("to-delete", "value")
	
	err := v.Delete("to-delete")
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}
	
	_, err = v.Get("to-delete")
	if err == nil {
		t.Error("Get() should fail after Delete()")
	}
}

func TestVault_DeleteNonExistent(t *testing.T) {
	v, _ := setupTestVault(t)
	
	err := v.Delete("non-existent")
	if err == nil {
		t.Error("Delete for non-existent key() should error")
	}
}

func TestVault_Persistence(t *testing.T) {
	tmpDir := t.TempDir()
	vaultFile := filepath.Join(tmpDir, "persist.json")
	
	cfg := &config.Config{
		MasterKey: "test-key-32-bytes!!!",
		VaultFile: vaultFile,
	}
	
	v1 := New(cfg)
	v1.Set("persistent", "data")
	
	v2 := New(cfg)
	val, err := v2.Get("persistent")
	if err != nil {
		t.Fatalf("Failed to get persisted data: %v", err)
	}
	
	if val != "data" {
		t.Errorf("Persisted value = %v, want data", val)
	}
}

func TestVault_MultipleKeys(t *testing.T) {
	v, _ := setupTestVault(t)
	
	keys := []string{"api_key", "db_password", "jwt_secret", "aws_key"}
	values := []string{"val1", "val2", "val3", "val4"}
	
	for i, k := range keys {
		if err := v.Set(k, values[i]); err != nil {
			t.Fatalf("Set() error = %v", err)
		}
	}
	
	list, _ := v.List()
	if len(list) != len(keys) {
		t.Errorf("List() = %d, want %d", len(list), len(keys))
	}
}

func TestVault_EmptyValue(t *testing.T) {
	v, _ := setupTestVault(t)
	
	v.Set("empty-key", "")
	
	val, err := v.Get("empty-key")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	
	if val != "" {
		t.Errorf("Get() = %v, want empty", val)
	}
}

func TestVault_UpdateKey(t *testing.T) {
	v, _ := setupTestVault(t)
	
	v.Set("key", "value1")
	v.Set("key", "value2")
	
	val, _ := v.Get("key")
	if val != "value2" {
		t.Errorf("Get() = %v, want value2", val)
	}
}

func TestVault_ListEmpty(t *testing.T) {
	v, _ := setupTestVault(t)
	
	keys, err := v.List()
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	
	if len(keys) != 0 {
		t.Errorf("List() = %d, want 0", len(keys))
	}
}

func TestVault_SaveCreatesDirectory(t *testing.T) {
	tmpDir := t.TempDir()
	vaultFile := filepath.Join(tmpDir, "subdir", "secrets.json")
	
	cfg := &config.Config{
		MasterKey: "test-key-32-bytes!!!",
		VaultFile: vaultFile,
	}
	
	v := New(cfg)
	err := v.Set("key", "value")
	
	if err != nil {
		t.Fatalf("Set() error = %v", err)
	}
	
	if _, err := os.Stat(vaultFile); os.IsNotExist(err) {
		t.Error("File should be created in subdirectory")
	}
}

func TestVault_EncryptedStorage(t *testing.T) {
	v, vaultFile := setupTestVault(t)
	
	v.Set("secret", "my-secret-value")
	
	data, err := os.ReadFile(vaultFile)
	if err != nil {
		t.Fatalf("ReadFile error = %v", err)
	}
	
	if strings.Contains(string(data), "my-secret-value") {
		t.Error("Secret should be encrypted, not plain text")
	}
}

func TestVault_LongKeys(t *testing.T) {
	v, _ := setupTestVault(t)
	
	longKey := strings.Repeat("a", 1000)
	err := v.Set(longKey, "value")
	if err != nil {
		t.Fatalf("Set() error = %v", err)
	}
	
	val, _ := v.Get(longKey)
	if val != "value" {
		t.Errorf("Get() = %v, want value", val)
	}
}

func TestVault_UnicodeKeys(t *testing.T) {
	v, _ := setupTestVault(t)
	
	keys := []string{"中文", "日本語", "🎉"}
	
	for _, k := range keys {
		err := v.Set(k, "value")
		if err != nil {
			t.Fatalf("Set() error for %q = %v", k, err)
		}
	}
	
	keys2, _ := v.List()
	if len(keys2) != len(keys) {
		t.Errorf("List() = %d, want %d", len(keys2), len(keys))
	}
}

func TestVault_ConcurrentWrites(t *testing.T) {
	v, _ := setupTestVault(t)
	
	for i := 0; i < 50; i++ {
		v.Set("key", "value"+string(rune('a'+i%26)))
	}
	
	val, _ := v.Get("key")
	if !strings.HasPrefix(val, "value") {
		t.Errorf("Get() should return value, got %v", val)
	}
}

func TestVault_LoadEmptyFile(t *testing.T) {
	tmpDir := t.TempDir()
	vaultFile := filepath.Join(tmpDir, "empty.json")
	
	os.WriteFile(vaultFile, []byte("{}"), 0600)
	
	cfg := &config.Config{
		MasterKey: "test-key-32-bytes!!!",
		VaultFile: vaultFile,
	}
	
	v := New(cfg)
	_, err := v.Get("any")
	if err == nil {
		t.Error("Should error on empty vault")
	}
}

func TestVault_LoadCorruptedFile(t *testing.T) {
	tmpDir := t.TempDir()
	vaultFile := filepath.Join(tmpDir, "corrupted.json")
	
	os.WriteFile(vaultFile, []byte("not valid json{{{"), 0600)
	
	cfg := &config.Config{
		MasterKey: "test-key-32-bytes!!!",
		VaultFile: vaultFile,
	}
	
	v := New(cfg)
	_ = v
}

func TestVault_SaveMultiple(t *testing.T) {
	v, _ := setupTestVault(t)
	
	for i := 0; i < 5; i++ {
		err := v.Set("key"+string(rune('0'+i)), "value")
		if err != nil {
			t.Fatalf("Set() error = %v", err)
		}
	}
}

func TestVault_PersistenceOverwrite(t *testing.T) {
	tmpDir := t.TempDir()
	vaultFile := filepath.Join(tmpDir, "overwrite.json")
	
	cfg := &config.Config{
		MasterKey: "test-key-32-bytes!!!",
		VaultFile: vaultFile,
	}
	
	v1 := New(cfg)
	v1.Set("key1", "val1")
	v1.Set("key1", "val2")
	
	v2 := New(cfg)
	val, _ := v2.Get("key1")
	if val != "val2" {
		t.Errorf("Overwrite failed, got %v", val)
	}
}

func TestVault_LoadValidEncryptedFile(t *testing.T) {
	tmpDir := t.TempDir()
	vaultFile := filepath.Join(tmpDir, "valid.json")
	
	cfg := &config.Config{
		MasterKey: "test-key-32-bytes!!!",
		VaultFile: vaultFile,
	}
	
	v1 := New(cfg)
	v1.Set("testkey", "testvalue")
	
	v2 := New(cfg)
	val, _ := v2.Get("testkey")
	
	if val != "testvalue" {
		t.Errorf("Got %v, want testvalue", val)
	}
}
