package main

import (
	"log"
	"mikit/internal/app"
	"os"
)

func main() {
	if err := app.NewAppCommand().Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
