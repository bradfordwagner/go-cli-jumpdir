package main

import (
	command_list "github.com/bradfordwagner/go-cli-jumpdir/internal/cmd/list"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list potential dirs to jump to",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		command_list.New().ListToStdOut()
		return
	},
}
