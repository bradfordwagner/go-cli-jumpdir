package main

import (
	command_list "github.com/bradfordwagner/go-cli-jumpdir/internal/cmd/list"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list potential dirs to jump to",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		res := command_list.New().List()
		for i, dir := range res {
			logrus.Infof("%d - %s", i, dir)
		}
		return
	},
}
