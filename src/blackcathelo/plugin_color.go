package main

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Color struct {
	Name     string
	RoleName string
	RID      string
}

var colors = []Color{
	{
		Name:     "粉紅色",
		RoleName: "Pink",
	},
	{
		Name:     "藍色",
		RoleName: "Blue",
	},
	{
		Name:     "紫色",
		RoleName: "purple",
	},
	{
		Name:     "變態色",
		RoleName: "violet",
	},
	{
		Name:     "紅色",
		RoleName: "red",
	},
	{
		Name:     "黃色",
		RoleName: "yellow",
	},
	{
		Name:     "橘色",
		RoleName: "orange",
	},
	{
		Name:     "灰色",
		RoleName: "grey",
	},
	{
		Name:     "深灰色",
		RoleName: "ash",
	},
}

var ColorIndex = 0

func ColorAddCond(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if strings.Index(m.Content, "我要變") == 0 {
		if colors[0].RID == "" {
			roles, err := s.GuildRoles(m.GuildID)
			if err != nil {
				s.ChannelMessageSend(
					m.ChannelID,
					"<@"+m.Author.ID+">"+"\n"+"我變不了顏色",
				)
				return false
			}

			for i, c := range colors {
				for _, r := range roles {
					if c.RoleName == r.Name {
						colors[i].RID = r.ID
					}
				}
			}
		}

		for i, c := range colors {
			if strings.Index(m.Content, c.Name) == (len(m.Content) - len(c.Name)) {
				ColorIndex = i
				return true
			}
		}
	}

	return false
}

func ColorAddGetResp(s *discordgo.Session, m *discordgo.MessageCreate) string {
	resp := "你已經變成" + colors[ColorIndex].Name + "了"
	err := s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, colors[ColorIndex].RID)
	if err != nil {
		log.Println(err)
		resp = "我變不了顏色"
	}
	ColorIndex = 0
	return resp
}

func ColorDelCond(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if strings.Index(m.Content, "我不要") == 0 {
		if colors[0].RID == "" {
			roles, err := s.GuildRoles(m.GuildID)
			if err != nil {
				s.ChannelMessageSend(
					m.ChannelID,
					"<@"+m.Author.ID+">"+"\n"+"我變不了顏色",
				)
				return false
			}

			for i, c := range colors {
				for _, r := range roles {
					if c.RoleName == r.Name {
						colors[i].RID = r.ID
					}
				}
			}
		}

		for i, c := range colors {
			if strings.Index(m.Content, c.Name) == (len(m.Content) - len(c.Name)) {
				ColorIndex = i
				return true
			}
		}
	}

	return false
}

func ColorDelGetResp(s *discordgo.Session, m *discordgo.MessageCreate) string {
	resp := "已經移除" + colors[ColorIndex].Name + "了"
	err := s.GuildMemberRoleRemove(m.GuildID, m.Author.ID, colors[ColorIndex].RID)
	if err != nil {
		log.Println(err)
		resp = "我變不了顏色"
	}
	ColorIndex = 0
	return resp
}
