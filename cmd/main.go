package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /alive", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Жив, цел, Орёл!")
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		//TODO: Логируем ошибку
	}
}
