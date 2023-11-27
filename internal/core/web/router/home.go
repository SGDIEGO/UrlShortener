package router

import (
	"github/SGDIEGO/UrlShortener/internal/core/web/handler"

	"github.com/go-chi/chi/v5"
)

type homeR struct {
	*Router
	hand *handler.Home
}

func HomeR(server *chi.Mux, url string) IRouter {
	var r homeR

	r.server = server
	r.url = url
	r.hand = handler.HomeH()

	return &r
}

func (r *homeR) load() error {
	server := *r.server

	server.Post(r.url, r.hand.PostUrl)

	return nil
}
