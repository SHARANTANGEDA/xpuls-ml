package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var GlobalCommandOption = struct {
	Debug bool
}{}

var rootCmd = &cobra.Command{
	Use:   "xpuls-ml-server",
	Short: "",
	Long:  "",
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&GlobalCommandOption.Debug, "debug", "d", false, "debug mode, output verbose output")
	rootCmd.AddCommand(getServeCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}
