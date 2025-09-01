package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/chmikata/incident-notification/internal/application"
	"github.com/spf13/cobra"
)

var apiKey string
var webhookUrl string
var template string

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "This is a command to notify you of an incident.",
	Long: `This is a command to notify you of an incident.

It collects incidents from Waroom and sends them to Slack.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		apiKey, _ = cmd.Flags().GetString("api-key")
		webhookUrl, _ = cmd.Flags().GetString("webhook-url")
		template, _ = cmd.Flags().GetString("template")
		if strings.HasPrefix(template, "@") {
			buf, err := os.ReadFile(strings.TrimPrefix(template, "@"))
			if err != nil {
				return fmt.Errorf("error retrieving template from path: %w", err)
			}
			template = string(buf)
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Executing action command...")
		logic := application.NewIncidentNotificatin(apiKey, webhookUrl)
		err := logic.Do(template)
		if err != nil {
			return fmt.Errorf("error executing action: %w", err)
		}
		fmt.Println("Action command executed successfully.")
		return nil
	},
}

func init() {
	actionCmd.Flags().StringP("api-key", "a", "", "Specify the API key for Waroom authentication")
	actionCmd.Flags().StringP("webhook-url", "w", "", "Specify the webhook URL for Slack")
	actionCmd.Flags().StringP("template", "t", "", "Specify a template for the output")

	actionCmd.MarkFlagRequired("api-key")
	actionCmd.MarkFlagRequired("webhook-url")
	actionCmd.MarkFlagRequired("template")

	rootCmd.AddCommand(actionCmd)
}
