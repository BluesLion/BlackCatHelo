package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ChoiceCond(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if (len(m.Content) > 6) &&
		strings.Index(strings.ToLower(m.Content[:6]), "choice") == 0 {
		return true
	}

	return false
}

func ChoiceGetResp(s *discordgo.Session, m *discordgo.MessageCreate) string {
	set := strings.Fields(m.Content)
	ret := PickOne(set[1:])

	resp := set[0] + " ["

	for i := 1; i < len(set); i++ {
		if i != 1 {
			resp = resp + ", "
		}
		resp = resp + set[i]
	}

	resp = resp + "]\n" + " → " + ret

	return resp
}
