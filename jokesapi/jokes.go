package jokesapi

import (
	"github.com/icelain/jokeapi"
	"log"
)

type JokesResp struct {
	Error    bool
	Category string
	JokeType string
	Joke     []string
	Flags    map[string]bool
	Id       float64
	Lang     string
}

func NewJokesResp() *JokesResp {
	return &JokesResp{}
}

func (j *JokesResp) GetJokes() {
	api := jokeapi.New()
	response, err := api.Fetch()
	if err != nil {
		log.Printf("Get jokes failed %v", err)
	}

	j.Id = response.Id
	j.Joke = response.Joke
	j.Lang = response.Lang
	j.Flags = response.Flags
	j.JokeType = response.JokeType
	j.Category = response.Category
}
