package main

import (
	"net/http"

	"github.com/RotemWald/smart-short-link/handlers"

	"github.com/RotemWald/smart-short-link/repositories"

	"github.com/RotemWald/smart-short-link/services"
)

func handleMultipleRoutes(handler func(http.ResponseWriter, *http.Request), routes ...string) {
	for _, route := range routes {
		http.HandleFunc(route, handler)
	}
}

func main() {
	repo := repositories.NewMemory()
	service := services.NewSmartUrl(repo)
	handler := handlers.NewSmartUrl(service)

	handleMultipleRoutes(handler.UUID, "/uuid", "/uuid/")
	handleMultipleRoutes(handler.Counter, "/counter", "/counter/")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
