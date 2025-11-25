package main

import (
	"PurpleHW/3-validation-api/configs"
	"PurpleHW/3-validation-api/internal/verify"
	"fmt"
	"net/http"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config: ", err)
	}
	router := &http.ServeMux{}
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{config})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	server.ListenAndServe()
}
