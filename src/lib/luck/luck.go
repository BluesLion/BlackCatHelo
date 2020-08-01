package luck

import (
	"math/rand"
	"time"
)

var results = []string{
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

func GetResults() string {
	rand.Seed(time.Now().UnixNano())
	return results[rand.Intn(len(results))]
}
