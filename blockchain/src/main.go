package main

import (
	"github.com/davecgh/go-spew/spew"
	"log"
	"time"
)

var blockchain Blockchain

func main() {

	go func() {
		t := time.Now()
		genesisBlock := Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		blockchain = append(blockchain, genesisBlock)
	}()
	log.Fatal(run())

}
