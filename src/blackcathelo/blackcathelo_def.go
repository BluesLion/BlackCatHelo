package main

import "github.com/bwmarrin/discordgo"

type Plugin struct {
	Condion       func(s *discordgo.Session, m *discordgo.MessageCreate) bool
	GetResp       func(s *discordgo.Session, m *discordgo.MessageCreate) string
	isRespMention bool
	isRespContent bool
	Help          string
}

var Events = []Plugin{
	{
		Condion:       LuckCond,
		GetResp:       LuckGetResp,
		isRespMention: true,
		isRespContent: true,
		Help:          "含有『運勢』關鍵字占卜 例如：出門運勢\n",
	},
	{
		Condion:       ChoiceCond,
		GetResp:       ChoiceGetResp,
		isRespMention: true,
		isRespContent: false,
		Help:          "choice <選擇1> <選擇2> <選擇3> ... 例如：choice A B C\n",
	},
	{
		Condion:       EatCond,
		GetResp:       EatGetResp,
		isRespMention: true,
		isRespContent: false,
		Help:          "吃什麼 如果不知道可以 吃什麼 <數字最大30> 例如：吃什麼 30\n",
	},
	{
		Condion:       TarotCond,
		GetResp:       TarotGetResp,
		isRespMention: true,
		isRespContent: true,
		Help:          "每日塔羅\n",
	},
	{
		Condion:       DndCond,
		GetResp:       DndGetResp,
		isRespMention: true,
		isRespContent: true,
		Help:          "九大陣營\n",
	},
	{
		Condion:       ColorAddCond,
		GetResp:       ColorAddGetResp,
		isRespMention: true,
		isRespContent: false,
		Help:          "我要變<顏色> 更換 ID 顏色\n\t目前可以變粉紅色,藍色,紫色,變態色,紅色,黃色,橘色,灰色,深灰色\n\t例如：我要變藍色\n",
	},
	{
		Condion:       ColorDelCond,
		GetResp:       ColorDelGetResp,
		isRespMention: true,
		isRespContent: false,
		Help:          "我不要<顏色> 更換 ID 顏色\n\t目前可以移除粉紅色,藍色,紫色,變態色,紅色,黃色,橘色,灰色,深灰色\n\t例如：我不要藍色\n",
	},
}
