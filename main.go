package main

import (
	"log"
	"net/http"
	"os"
	"gohan/app"
)

func main() {
	app := app.New()

	http.HandlerFunc("/", app.Router.ServeHTTP)

	err := http.ListenAndServe(":9000", nil)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}