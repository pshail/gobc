package main

import (
	"fmt"
	"log"
	"os"

	"github.com/heckdevice/gobc/core"
	"github.com/heckdevice/gobc/webservice"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	testRun := os.Getenv("TEST_RUN")
	if testRun == "true" {
		testAdd()
	} else {
		initGenesisBlock()
		log.Fatal(webservice.Run())
	}
}

func initGenesisBlock() {
	//Initialize the genesis block
	go func() {
		core.Init()
	}()
}
func testAdd() {
	for i := 0; i < 10000; i++ {
		_, _, err := core.Add(map[string]interface{}{"DataPoint": i})
		if err != nil {
			log.Fatal(err)
		}
	}
	bcData, err := core.GetChain()
	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println(*bcData)
	}
}
