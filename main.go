package main

import (
	"context"
	"embed"
	"errors"
	"os"

	"log"

	"github.com/cf-demo-company/receipt-app/server"
)

// content holds our web server templates
//go:embed template/*
//go:embed static/dist/*
var content embed.FS

func main() {
	err := run()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}

// run is the actual entrypoint of the program
// and returns an error if anything fails
func run() error {
	// We store receipts in an S3 bucket
	// try and load config relating to this
	bucketName := os.Getenv("S3_BUCKET_NAME")
	if bucketName == "" {
		return errors.New("S3_BUCKET_NAME variable was not provided")
	}

	log.Default().Println("starting Receipt App server...")

	storage, err := server.NewStorage(context.Background(), bucketName)
	if err != nil {
		return err
	}

	s := server.NewServer(storage, content)
	return s.Run()
}
