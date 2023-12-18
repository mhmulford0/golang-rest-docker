package routes

import (
	"net/http"

	"github.com/mhmulford0/golang-rest-docker/internal/handlers"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)

	return mux

}
