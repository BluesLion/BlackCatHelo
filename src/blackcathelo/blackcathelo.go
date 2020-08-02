package main

import (
	"fmt"
	"lib/dndalign"
	"lib/eat"
	"lib/luck"
	"lib/rcore"
	"lib/tarot"
	"log"
	"os"
	"os/signal"
	"strconv"
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

	respHeader := "<@" + m.Author.ID + ">"

	if strings.Index(m.Content, "運勢") >= 0 {
		s.ChannelMessageSend(
			m.ChannelID,
			respHeader+"\n"+m.Content+"："+luck.GetResults(),
		)
		return
	}

	if (len(m.Content) > 6) &&
		strings.Index(strings.ToLower(m.Content[:6]), "choice") == 0 {
		set := strings.Fields(m.Content)
		ret := rcore.PickOne(set)

		resp := respHeader + "\n" + set[0] + " ["

		for i := 1; i < len(set); i++ {
			if i != 1 {
				resp = resp + ", "
			}
			resp = resp + set[i]
		}

		resp = resp + "]\n" + " → " + ret
		s.ChannelMessageSend(
			m.ChannelID,
			resp,
		)
		return
	}

	if strings.Index(m.Content, "吃什麼") == 0 {
		set := strings.Fields(m.Content)
		resp := eat.GetResults()
		title := ""

		if len(set) == 2 {
			count, err := strconv.Atoi(set[1])
			if err != nil {
				count = 1
			}

			if count > 30 {
				count = 30
			}

			title = fmt.Sprintf("%d 個吃什麼 → ", count)

			for i := 1; i < count; i++ {
				ret := eat.GetResults()
				if strings.Index(resp, ret) < 0 {
					resp = resp + ", "
					resp = resp + ret
				} else {
					i--
				}
			}
		}

		s.ChannelMessageSend(
			m.ChannelID,
			respHeader+"\n"+title+resp,
		)
		return
	}

	switch m.Content {
	case "九大陣營":
		s.ChannelMessageSend(
			m.ChannelID,
			respHeader+"\n"+m.Content+" → "+dndalign.GetResults(),
		)
	case "每日塔羅":
		s.ChannelMessageSend(
			m.ChannelID,
			respHeader+"\n"+m.Content+" → "+tarot.GetResults(),
		)
	case "我要變藍色":
		r, err := s.GuildRoles(m.GuildID)
		if err != nil {
			s.ChannelMessageSend(
				m.ChannelID,
				respHeader+"\n"+"我變不了顏色",
			)
			return
		}

		for i := range r {
			if r[i].Name == "Blue" {
				s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, r[i].ID)
				s.ChannelMessageSend(
					m.ChannelID,
					respHeader+"\n"+"OK",
				)
				return
			}
		}
	case "我不要藍色":
		r, err := s.GuildRoles(m.GuildID)
		if err != nil {
			s.ChannelMessageSend(
				m.ChannelID,
				respHeader+"\n"+"我變不了顏色",
			)
			return
		}

		for i := range r {
			if r[i].Name == "Blue" {
				s.GuildMemberRoleRemove(m.GuildID, m.Author.ID, r[i].ID)
				s.ChannelMessageSend(
					m.ChannelID,
					respHeader+"\n"+"OK",
				)
				return
			}
		}
	case "海螺幫幫我":
		s.ChannelMessageSend(
			m.ChannelID,
			respHeader+"\n"+Help,
		)
	default:
		return
	}
}
