package routes

import (
	"net/http"
)

func Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	return mux

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	w.Write([]byte("StatusOK"))
}
