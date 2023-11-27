package router

import "github.com/go-chi/chi/v5"

type Router struct {
	server *chi.Mux
	url    string
}

type IRouter interface {
	load() error
}

func LoadRouters(server *chi.Mux) error {
	// Routers
	HOME_ROUTER := HomeR(server, "/")

	// Load home router
	err := HOME_ROUTER.load()
	if err != nil {
		return err
	}

	return nil
}
