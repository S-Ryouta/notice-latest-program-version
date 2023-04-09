package notification

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func SendNotificationToDiscord(language string, version string) error {
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

	message := fmt.Sprintf("%sの最新バージョンがリリースされました: %s", language, version)
	channelID := os.Getenv("DISCORD_CHANNEL_ID")

	_, err = dg.ChannelMessageSend(channelID, message)
	return err
}
