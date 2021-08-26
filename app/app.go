package app

import "github.com/gorilla/mux"

type App struck {
	Router *mux.Router
}

func New() *App {
	return &App{
			Router: mux.NewRouter()
	}
}