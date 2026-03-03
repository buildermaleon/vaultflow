package vault

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dablon/vaultflow/internal/config"
	"github.com/dablon/vaultflow/internal/crypto"
)

type Vault struct {
	config  *config.Config
	secrets map[string]string
	enc    *crypto.Encryptor
}

func New(cfg *config.Config) *Vault {
	v := &Vault{
		config:  cfg,
		secrets: make(map[string]string),
		enc:    crypto.New(cfg.MasterKey),
	}
	v.load() // Load existing secrets
	return v
}

func (v *Vault) Set(key, value string) error {
	encrypted, err := v.enc.Encrypt(value)
	if err != nil {
		return fmt.Errorf("encryption failed: %w", err)
	}
	v.secrets[key] = encrypted
	return v.save()
}

func (v *Vault) Get(key string) (string, error) {
	val, ok := v.secrets[key]
	if !ok {
		return "", fmt.Errorf("secret not found: %s", key)
	}
	return v.enc.Decrypt(val)
}

func (v *Vault) List() ([]string, error) {
	keys := make([]string, 0, len(v.secrets))
	for k := range v.secrets {
		keys = append(keys, k)
	}
	return keys, nil
}

func (v *Vault) Delete(key string) error {
	if _, ok := v.secrets[key]; !ok {
		return fmt.Errorf("secret not found: %s", key)
	}
	delete(v.secrets, key)
	return v.save()
}

func (v *Vault) save() error {
	if err := os.MkdirAll(filepath.Dir(v.config.VaultFile), 0700); err != nil {
		return err
	}
	data, err := json.Marshal(v.secrets)
	if err != nil {
		return err
	}
	return os.WriteFile(v.config.VaultFile, data, 0600)
}

func (v *Vault) load() error {
	data, err := os.ReadFile(v.config.VaultFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No secrets file yet
		}
		return err
	}
	return json.Unmarshal(data, &v.secrets)
}
