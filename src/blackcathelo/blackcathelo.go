package main

import (
	"lib/dndalign"
	"lib/luck"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	Token := getToken(".env")
	if Token == "" {
		log.Println("No token provided. Please add .env file")
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	respHeader := "<@" + m.Author.ID + ">" + "\n" + m.Content

	if strings.Index(m.Content, "運勢") >= 0 {
		s.ChannelMessageSend(
			m.ChannelID,
			respHeader+"："+luck.GetResults(),
		)
		return
	}

	switch m.Content {
	case "九大陣營":
		s.ChannelMessageSend(
			m.ChannelID,
			respHeader+" → "+dndalign.GetResults(),
		)
	default:
		return
	}
}
