package main

import (
	"fmt"
	"log"

	"github.com/heckdevice/gobc/core"
)

func main() {
	testAdd()
}

func testAdd() {
	for i := 0; i < 10000; i++ {
		_, err := core.Add(map[string]interface{}{"DataPoint": i})
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
