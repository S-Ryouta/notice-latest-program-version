package notification

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func SendNotificationToDiscord(version string) error {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}

	err = dg.Open()
	if err != nil {
		return err
	}
	defer dg.Close()

	message := fmt.Sprintf("Golangの最新バージョンがリリースされました: %s", version)
	channelID := os.Getenv("DISCORD_CHANNEL_ID")

	_, err = dg.ChannelMessageSend(channelID, message)
	return err
}
