package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var luckResults = []string{
	"超吉",
	"超級上吉",
	"大吉",
	"吉",
	"中吉",
	"小吉",
	"吉",
	"小吉",
	"吉",
	"吉",
	"中吉",
	"吉",
	"中吉",
	"吉",
	"中吉",
	"小吉",
	"末吉",
	"吉",
	"中吉",
	"小吉",
	"末吉",
	"中吉",
	"小吉",
	"小吉",
	"吉",
	"小吉",
	"末吉",
	"中吉",
	"小吉",
	"凶",
	"小凶",
	"沒凶",
	"大凶",
	"很凶",
	"你不要知道比較好呢",
	"命運在手中,何必問我",
}

func LuckCond(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if (strings.Index(m.Content, "運勢") == 0) ||
		(strings.Index(m.Content, "運勢") == (len(m.Content) - len("運勢"))) {
		return true
	}

	return false
}

func LuckGetResp(s *discordgo.Session, m *discordgo.MessageCreate) string {
	return "：" + PickOne(luckResults)
}
