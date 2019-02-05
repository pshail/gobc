package core

import (
	"time"
)

var (
	// Blockchain - Simple chain of blocks
	Blockchain BC

	genesisHash     = "dd102105cdd3a38ce77562d856cfce6d"
	genesisSeedData = map[string]interface{}{"whoami": 42}
)

// BC a.k.a Blockchain, a simple chain of blocks
type BC struct {
	// Chain - Blocks in BC
	Chain []Block `json:"chain"`
}

func getGenesisBlock() Block {
	now := time.Now()
	genTimeStamp := now.String()
	block := Block{0, &genTimeStamp, nil, &genesisHash, nil}
	return block
}

func init() {
	genBlock := getGenesisBlock()
	Blockchain.Chain = append(Blockchain.Chain, genBlock)
}

/******* Core functions of the chain **********/
