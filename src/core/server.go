package core

import (
	"drafter/src/config"
	"net/http"
)

func InitServer() *http.Server {
	serverPort, err := config.GetEnvVar("PORT")
	if err != nil {
		serverPort = "8788"
	}
	hostname, err := config.GetEnvVar("HOST")
	if err != nil {
		hostname = "api.drafter.local"
	}

	router := InitRouting()

	serverAddress := hostname + ":" + serverPort

	server := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  config.DefaultReadTimeout,
		WriteTimeout: config.DefaultWriteTimeout,
		Handler:      router,
	}

	return server
}
