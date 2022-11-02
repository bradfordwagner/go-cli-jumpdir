package main

import (
	command_add "github.com/bradfordwagner/go-cli-jumpdir/internal/cmd/add"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a dir to the jumpdir backend",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return command_add.New().AddDirectories(args)
	},
}
