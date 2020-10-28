package cmd

import (
	"fmt"
	"github.com/Shitovdm/go-file-api/conf"
	"github.com/Shitovdm/go-ping/Ping"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "Use mode for work",
	Long: `Use mode api for work
		  config file can be selected by --config=filename key`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initAll)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}

func initAll() {
	conf.ReadConfigFiles("./etc/", cfgFile)
	go Ping.Serve(conf.GetPingApiPort())
}
