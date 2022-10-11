package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	app "Abylkaiyr/groupie-tracker/app"
)

func main() {
	addr := flag.String("port", "8181", "Network port")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &app.AppServer{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,

		Addr:     ":" + *addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	app.InfoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
