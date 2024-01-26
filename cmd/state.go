package cmd

import (
	"github.com/hardcore-os/plato/state"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand()
}

var stateCmd = &cobra.Command{
	Use: "state",
	Run: StateHandle,
}

func StateHandle(cmd *cobra.Command, args []string) {
	state.RunMain(ConfigPath)
}
