package dndalign

import "lib/rcore"

var results = []string{
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

func GetResults() string {
	return rcore.PickOne(results)
}
