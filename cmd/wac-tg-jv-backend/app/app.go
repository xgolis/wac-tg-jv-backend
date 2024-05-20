package app

import (
	"net/http"
	"time"
)

type App struct {
	Server *http.Server
}

func NewApp() *App {
	mux := MakeHandlers()

	return &App{
		Server: &http.Server{
			Addr:           "0.0.0.0:8080",
			Handler:        mux,
			ReadTimeout:    100 * time.Second,
			WriteTimeout:   100 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}
