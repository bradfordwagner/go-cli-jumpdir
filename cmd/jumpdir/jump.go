package main

import (
	command_add "github.com/bradfordwagner/go-cli-jumpdir/internal/cmd/add"
	command_list "github.com/bradfordwagner/go-cli-jumpdir/internal/cmd/list"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(jumpCmd)
}

var jumpCmd = &cobra.Command{
	Use:   "jump",
	Short: "jumps to dir",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		res := command_list.New().List()
		return res, cobra.ShellCompDirectiveDefault
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return command_add.New().AddDirectories(args)
	},
}
