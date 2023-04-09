package notification

import (
	"fmt"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func SendNotificationToLine(version string) error {
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	channelToken := os.Getenv("LINE_CHANNEL_TOKEN")

	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("Golangの最新バージョンがリリースされました: %s", version)
	if _, err := bot.BroadcastMessage(linebot.NewTextMessage(message)).Do(); err != nil {
		return err
	}

	return nil
}
