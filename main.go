package main

import (
	"app_backend/internal"
	"app_backend/internal/config"
	"log"
)

func main() {
	config.Load()
	if err := internal.StartServer(); err != nil {
		log.Fatal(err)
	}
}
