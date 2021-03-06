package cmd

import (
	"github.com/bionic-dev/bionic/database"
	"github.com/bionic-dev/bionic/imports"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset [provider]",
	Short: "Reset provider data stored in local db",
	RunE: func(cmd *cobra.Command, args []string) error {
		providerName := args[0]

		db, err := database.New(dbPath)
		if err != nil {
			return err
		}

		manager, err := imports.NewManager(db, imports.DefaultProviders(db))
		if err != nil {
			return err
		}

		provider, err := manager.GetByName(providerName)
		if err != nil {
			return err
		}

		return manager.Reset(provider)
	},
	Args: cobra.MinimumNArgs(1),
}
