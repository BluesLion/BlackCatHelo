package main

import (
	"io/ioutil"
	"log"
	"strings"
)

const Help = "1. 含有『運勢』關鍵字占卜 例如：出門運勢\n" +
	"2. choice <選擇1> <選擇2> <選擇3> ... 例如：choice A B C\n" +
	"3. 吃什麼 如果不知道可以 吃什麼 <數字最大30> 例如：吃什麼 30\n" +
	"4. 每日塔羅\n" +
	"5. 九大陣營\n"

/*
 * Attempt the setting file only have one line:
 * DISCORD_CHANNEL_SECRET=<TOKEN>
 */
func getToken(localtion string) string {
	content, err := ioutil.ReadFile(localtion)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	line := strings.Split(string(content), "=")
	if line[0] == "DISCORD_CHANNEL_SECRET" {
		return line[1]
	}
	return ""
}
