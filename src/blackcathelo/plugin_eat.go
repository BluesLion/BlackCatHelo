package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var EatResults = []string{
	"飯",
	"炒飯",
	"燉飯",
	"粥",
	"便當",
	"蓋飯",
	"自助餐",
	"麵",
	"鍋燒麵",
	"湯麵",
	"炒麵",
	"義大利麵",
	"吃屎",
	"牛肉麵",
	"麵包",
	"麥當勞",
	"肯德基",
	"摩斯",
	"丹丹",
	"頂呱呱",
	"牛排",
	"披薩",
	"很貴的壽司",
	"便宜貨壽司",
	"三明治",
	"小籠包",
	"快炒",
	"吃土",
	"壽喜燒",
	"燒肉",
	"小火鍋",
	"貴火鍋",
	"水餃",
	"煎餃",
	"饅頭",
	"烤地瓜",
	"肉包",
	"焗烤",
	"永和豆漿",
	"叉燒飯",
	"燒鴨飯",
	"三寶飯",
	"燒肉飯",
	"油雞飯",
	"滷味",
	"滷肉飯",
	"陽春麵",
	"餛飩麵",
	"7-11",
	"全家",
	"燒烤",
	"Subway",
	"素食",
	"肉粽",
	"蛋餅",
	"拉麵",
	"絕食(抗議)",
	"挫冰",
	"涼麵",
	"綠豆湯",
	"雞排",
	"鹹酥雞",
	"咖哩飯",
	"麻辣燙",
	"麻辣鴨血",
	"北京烤鴨",
	"夜市",
	"車輪餅",
	"雞蛋糕",
	"肉圓",
	"抓餅",
	"碗粿",
	"冰淇淋",
	"鐵板燒",
	"臭豆腐",
	"粉圓豆花",
	"粉圓冰",
}

func EatCond(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if strings.Index(m.Content, "吃什麼") == 0 {
		return true
	}

	return false
}

func EatGetResp(s *discordgo.Session, m *discordgo.MessageCreate) string {
	set := strings.Fields(m.Content)
	resp := PickOne(EatResults)
	title := ""

	if len(set) == 2 {
		count, err := strconv.Atoi(set[1])
		if err != nil {
			count = 1
		}

		if count > 30 {
			count = 30
		}

		title = fmt.Sprintf("%d/%d 個吃什麼 → ", count, len(EatResults))

		for i := 1; i < count; i++ {
			ret := PickOne(EatResults)
			if strings.Index(resp, ret) < 0 {
				resp = resp + ", "
				resp = resp + ret
			} else {
				i--
			}
		}
	}

	return title + resp
}
