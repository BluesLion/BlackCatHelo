package main

import (
	"io/ioutil"
	"log"
	"strings"
)

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
