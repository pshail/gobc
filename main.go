package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/heckdevice/gobc/core"
)

func main() {
	testAdd()
}

func testAdd() {
	for i := 0; i < 2; i++ {
		_, err := core.Add(map[string]interface{}{"DataPoint": i})
		if err != nil {
			log.Fatal(err)
		}
	}
	bcData, err := core.GetChain()
	if err != nil {
		log.Fatal(err)

	} else {
		spew.Dump(bcData)
	}
}
