package main

import (
	"PurpleHW/3-validation-api/configs"
	"PurpleHW/3-validation-api/internal/verify"
	"fmt"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	router := http.NewServeMux()
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{config})
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("Listening on port 8080")

	server.ListenAndServe()

}
