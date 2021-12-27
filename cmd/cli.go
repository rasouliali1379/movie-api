package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	addCommands()
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
