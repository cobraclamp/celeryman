package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cobraclamp/celeryman/config"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)
	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}

		if m.Content == "!oyster" {

			embed := discordgo.MessageEmbed{
				Title: "Now, thats a nice smile.",
				Image: &discordgo.MessageEmbedImage{
					URL: "http://ericsizer.com/oyster.png",
				},
			}

			msg, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)

			if err != nil {
				println(err.Error())
			}
			println(msg)
		}
	}
}
