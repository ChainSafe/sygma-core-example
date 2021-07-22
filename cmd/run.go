package cmd

import (
	"github.com/ChainSafe/chainbridge-core-example/example"
	"github.com/ChainSafe/chainbridge-core/config"
	"github.com/spf13/cobra"
)

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run example app",
		Long:  "Run example app",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := example.Run(); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	//rootCMD.AddCommand(evmClient.CLI()) // Example of how CLI should be registered
	rootCmd.AddCommand(runCmd)
	config.BindFlags(runCmd)
}
