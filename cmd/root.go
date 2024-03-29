// Package cmd is the package that contains all of the commands for the application.
package cmd

import (
	"os"

	"github.com/Serpentiel/go-cli/internal/provide"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// rootCmd is the rootCmd command for the application.
var rootCmd = &cobra.Command{
	Use: "go-cli",
	Run: func(cmd *cobra.Command, _ []string) {
		fx.New(
			fx.Supply(cmd),

			fx.Provide(
				provide.Viper,
				provide.Logger,
			),

			fx.WithLogger(provide.FxeventLogger),
		).Run()
	},
}

// Execute is the function that is called to execute the rootCmd command.
func Execute() {
	rootCmd.PersistentFlags().StringP("config", "c", "", "defines the config file to use")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
