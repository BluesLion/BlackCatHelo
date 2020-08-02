package rcore

import (
	"math/rand"
	"time"
)

func PickOne(set []string) string {
	rand.Seed(time.Now().UnixNano())
	return set[rand.Intn(len(set))]
}
