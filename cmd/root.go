package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Short: "Simple archiver",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		handleErr(err)
	}

}

func handleErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
