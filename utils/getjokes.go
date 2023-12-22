package utils

import (
	"math/rand"
	"time"

	"Graph-QL-Jokes/model"
)

func GetRandomJoke() *model.Joke {
	var jokes []model.Joke
	db.Find(&jokes)

	if len(jokes) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(jokes))
	return &jokes[index]
}
