package main

import (
	"log"

	"github.com/heckdevice/gobc/core"
	"github.com/heckdevice/gobc/webservice"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	initGenesisBlock()
	log.Fatal(webservice.Run())

}

func initGenesisBlock() {
	//Initialize the genesis block
	go func() {
		core.Init()
	}()
}
