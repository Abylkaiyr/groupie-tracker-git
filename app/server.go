package app

import (
	"log"
	"net/http"
)

type AppServer struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func (app *AppServer) Routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("ui"))
	mux.Handle("/ui/", http.StripPrefix("/ui", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/detail", app.details)
	mux.HandleFunc("/filters", app.filters)
	mux.HandleFunc("/search", app.search)
	return mux
}
