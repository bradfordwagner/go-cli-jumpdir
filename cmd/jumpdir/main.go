package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "jumpdir",
}

func main() {
	// setup logrus
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// cobra
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
