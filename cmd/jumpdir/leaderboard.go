package main

import (
	command_leaderboard "github.com/bradfordwagner/go-cli-jumpdir/internal/cmd/leaderboard"
	"github.com/spf13/cobra"
)

var (
	leaderboardReverse bool
	leaderboardLimit   int
)

func init() {
	rootCmd.AddCommand(leaderboardCmd)
	leaderboardCmd.Flags().BoolVarP(&leaderboardReverse, "reverse", "r", false, "reverse sort order (ascending)")
	leaderboardCmd.Flags().IntVarP(&leaderboardLimit, "number", "n", 0, "limit number of results")
}

var leaderboardCmd = &cobra.Command{
	Use:   "leaderboard",
	Short: "show directories sorted by weight",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		command_leaderboard.New(leaderboardReverse, leaderboardLimit).LeaderboardToStdOut()
		return
	},
}
