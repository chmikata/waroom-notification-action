package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "incident-notification",
	Short: "This is a command to notify you of an incident.",
	Long: `This is a command to notify you of an incident.

It collects incidents from Waroom and sends them to Slack.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
