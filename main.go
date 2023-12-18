package main

import (
	"fmt"
	"net/http"

	"github.com/mhmulford0/golang-rest-docker/internal/routes"
)

func main() {
	router := routes.Router()

	fmt.Printf("Listening on port 8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic("Could not start the server")
	}
}
