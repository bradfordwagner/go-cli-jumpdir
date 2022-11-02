package main

import (
	command_delete "github.com/bradfordwagner/go-cli-jumpdir/internal/cmd/delete"
	command_list "github.com/bradfordwagner/go-cli-jumpdir/internal/cmd/list"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes dirs to the jumpdir backend",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		res := command_list.New().List()
		return res, cobra.ShellCompDirectiveDefault
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return command_delete.New().DeleteDirectories(args)
	},
}
