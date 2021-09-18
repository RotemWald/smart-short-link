package main

import (
	"net/http"

	"github.com/RotemWald/smart-short-link/handlers"

	"github.com/RotemWald/smart-short-link/repositories"

	"github.com/RotemWald/smart-short-link/services"
)

func handleMultipleEndpoints(endpoints []string, handler func(http.ResponseWriter, *http.Request)) {
	for _, endpoint := range endpoints {
		http.HandleFunc(endpoint, handler)
	}
}

func main() {
	repo := repositories.NewMemory()
	service := services.NewSmartUrl(repo)
	handler := handlers.NewSmartUrl(service)

	handleMultipleEndpoints([]string{"/uuid", "/uuid/"}, handler.UUID)
	handleMultipleEndpoints([]string{"/counter", "/counter/"}, handler.Counter)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
