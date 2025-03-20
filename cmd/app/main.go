package main

import (
	"basicblockchain/internal/api"
	"basicblockchain/internal/model"
	"log"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	rootDir, err := filepath.Abs("../../")
	if err != nil {
		log.Fatalf("Error getting root directory: %v", err)
	}

	// Load the .env file from the root directory
	envPath := filepath.Join(rootDir, ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	blockChain := model.NewBlockChain()

	go func() {
		spew.Dump(blockChain)
	}()

	log.Fatal(api.RunServer(blockChain))
}
