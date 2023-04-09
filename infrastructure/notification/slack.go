package notification

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func SendNotificationToSlack(version string) error {
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)

	message := fmt.Sprintf("Golangの最新バージョンがリリースされました: %s", version)
	channelID := os.Getenv("SLACK_CHANNEL_ID")

	_, _, err := api.PostMessage(channelID, slack.MsgOptionText(message, true))
	if err != nil {
		return err
	}

	return nil
}
