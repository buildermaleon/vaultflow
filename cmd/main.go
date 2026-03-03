package main

import (
	"fmt"
	"os"

	"github.com/dablon/vaultflow/internal/vault"
	"github.com/dablon/vaultflow/internal/config"
	"github.com/spf13/cobra"
)

var version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:     "vaultflow",
	Short:   "Secrets Management CLI",
	Long:    "VaultFlow - Secure secrets management with AES-256-GCM encryption",
	Version: version,
}

var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Store a secret",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		v := vault.New(config.Default())
		if err := v.Set(args[0], args[1]); err != nil {
			return err
		}
		fmt.Println("✓ Secret stored securely")
		return nil
	},
}

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Retrieve a secret",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		v := vault.New(config.Default())
		val, err := v.Get(args[0])
		if err != nil {
			return err
		}
		fmt.Println(val)
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all secrets",
	RunE: func(cmd *cobra.Command, args []string) error {
		v := vault.New(config.Default())
		keys, err := v.List()
		if err != nil {
			return err
		}
		if len(keys) == 0 {
			fmt.Println("No secrets stored")
			return nil
		}
		fmt.Println("Stored secrets:")
		for _, k := range keys {
			fmt.Printf("  • %s\n", k)
		}
		return nil
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [key]",
	Short: "Delete a secret",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		v := vault.New(config.Default())
		if err := v.Delete(args[0]); err != nil {
			return err
		}
		fmt.Println("✓ Secret deleted")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(setCmd, getCmd, listCmd, deleteCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
