package main

import (
	"log"
	"net/http"
	"os"

	"ports-server/configs"
	adapter "ports-server/internal/adapter/handler/http"
	repository "ports-server/internal/adapter/repository/ports"
	"ports-server/internal/core/util"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c, err := configs.NewConfig()
	if err != nil {
		return err
	}

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	l := log.New(file, "TEST_PORTS: ", log.LstdFlags)

	storagePorts := repository.NewStorageIn(c)

	serverHandler := adapter.New(l, storagePorts)
	http.Handle("/read", util.RateLimiter(l, serverHandler.Read))
	http.Handle("/write", util.RateLimiter(l, serverHandler.Write))
	log.Println("Server listening on port :8080")
	err = http.ListenAndServe(c.Http.Port, nil)
	if err != nil {
		log.Println("There was an error listening on port :8080", err)
	}

	return nil
}
