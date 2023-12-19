package handler

import (
	"context"
	"go-templ/jokesapi"
	"go-templ/view"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	newJoke := jokesapi.NewJokesResp()
	newJoke.GetJokes()

	err := view.Base(view.Joke(newJoke.Joke)).Render(context.Background(), w)
	if err != nil {
		log.Printf("Rendering home failed %v", err)
	}

}


