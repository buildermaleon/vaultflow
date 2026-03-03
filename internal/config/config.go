package config

import "os"

type Config struct {
	MasterKey string
	VaultFile string
}

func Default() *Config {
	return &Config{
		MasterKey: os.Getenv("VAULT_MASTER_KEY"),
		VaultFile: os.ExpandEnv("$HOME/.vaultflow/secrets.json"),
	}
}
