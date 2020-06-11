package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"log"
	"time"
)

var blockchain Blockchain

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		blockchain = append(blockchain, genesisBlock)
	}()
	log.Fatal(run())

}
