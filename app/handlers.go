package app 

import (
	"fmt"
	"net/http"
)

func(a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "gohan")
	}
}