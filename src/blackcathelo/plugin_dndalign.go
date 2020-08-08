package main

import "github.com/bwmarrin/discordgo"

var DndResults = []string{
	"守序善良",
	"守序中立",
	"守序邪惡",
	"中立善良",
	"絕對中立",
	"中立邪惡",
	"混亂善良",
	"混亂中立",
	"混亂邪惡",
}

func DndCond(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if m.Content == "九大陣營" {
		return true
	}

	return false
}

func DndGetResp(s *discordgo.Session, m *discordgo.MessageCreate) string {
	return " → " + PickOne(DndResults)
}
