package cmd

import (
	"csf/cmd/api"
	"csf/cmd/apis"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "go-admin",
	Short:        "go-admin",
	SilenceUsage: true,
	Long:         `go-admin`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
	},
}
var configYml string

func init() {
	rootCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/configs.yml", "Start server with provided configuration file")

	// 初始化api模块
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(apis.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
