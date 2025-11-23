package main

import (
	"PurpleHW/3-validation-api/configs"
	"PurpleHW/3-validation-api/internal/verify"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	router := &http.ServeMux{}
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{config})
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
