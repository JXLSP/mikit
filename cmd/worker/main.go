package main

import (
	"log"
	"mikit/internal/worker"
	"os"
)

func main() {
	if err := worker.NewWorkerCommand().Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
