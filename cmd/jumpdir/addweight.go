package main

import (
	command_addweight "github.com/bradfordwagner/go-cli-jumpdir/internal/cmd/addweight"
	command_list "github.com/bradfordwagner/go-cli-jumpdir/internal/cmd/list"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addWeightCmd)
}

var addWeightCmd = &cobra.Command{
	Use:   "addweight",
	Short: "adds weight to dir",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		res := command_list.New().List()
		return res, cobra.ShellCompDirectiveDefault
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return command_addweight.New().AddWeight(args[0])
	},
}
