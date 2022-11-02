package main

import (
	command_delete "github.com/bradfordwagner/go-cli-jumpdir/internal/cmd/delete"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes dirs to the jumpdir backend",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return command_delete.New().DeleteDirectories(args)
	},
}
