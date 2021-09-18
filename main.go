package main

import (
	"net/http"

	"github.com/RotemWald/smart-short-link/handlers"

	"github.com/RotemWald/smart-short-link/repositories"

	"github.com/RotemWald/smart-short-link/services"
)

func handleMultipleEndpoints(handler func(http.ResponseWriter, *http.Request), endpoints ...string) {
	for _, endpoint := range endpoints {
		http.HandleFunc(endpoint, handler)
	}
}

func main() {
	repo := repositories.NewMemory()
	service := services.NewSmartUrl(repo)
	handler := handlers.NewSmartUrl(service)

	handleMultipleEndpoints(handler.UUID, "/uuid", "/uuid/")
	handleMultipleEndpoints(handler.Counter, "/counter", "/counter/")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
