package cmd

import (
	"github.com/Shitovdm/go-file-api/conf"
	"github.com/Shitovdm/go-file-api/src/api"
	"github.com/Shitovdm/go-log/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmd)
}

var cmd = &cobra.Command{
	Use:   "api",
	Short: "Runs API services",
	Long:  `Runs HTTP interlayer API.`,
	Run: func(cmd *cobra.Command, args []string) {
		loggerInstance := logger.NewLoggerInstance()

		loggerInstance.Info("Running HTTP API server...", conf.GetLogCategory())
		_ = api.NewServer(conf.GetFsRootDir()).StartServe(conf.GetFileApiPort())
	},
}
