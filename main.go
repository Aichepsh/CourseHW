package main

import (
	"PurpleHW/1-concurrency"
	RandomAPI "PurpleHW/2-RandomAPI"
	"net/http"
)

func main() {
	concurrency.Conc()
	router := http.NewServeMux()
	RandomAPI.NewRandomAPI(router)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
