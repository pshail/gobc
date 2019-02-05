package main

import (
	"log"

	"github.com/heckdevice/gobc/core"
)

func main() {
	for i := 0; i < 100; i++ {
		_, err := core.Add(map[string]interface{}{"DataPoint": i})
		if err != nil {
			log.Fatal(err)
		}
	}
}
