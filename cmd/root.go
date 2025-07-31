package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vessl",
	Short: "Vessl is a tool for managing docker containers",
	Long: `Vessl is a tool for managing docker containers.
	It allows you to create, start, stop, and remove docker containers.
	It also allows you to list and inspect docker containers.
	`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
