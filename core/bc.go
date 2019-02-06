package core

import (
	"encoding/json"
	"sync"
	"time"
)

var (
	// Blockchain - Simple chain of blocks
	blockchain      = &BC{}
	mutex           = &sync.Mutex{}
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

// init - Initialize a Blockchain with genesis block
func init() {
	mutex.Lock()
	if blockchain.Chain == nil || len(blockchain.Chain) == 0 {
		blockchain.Chain = make([]Block, 1, 1)
		genBlock := getGenesisBlock()
		blockchain.Chain[0] = genBlock
	}
	mutex.Unlock()
}

// Add - adds new data to the blockchain
// functions critical sections are mutexed
func Add(data interface{}) (bool, error) {
	newBlock, err := blockchain.GenerateBlock(data)
	if err != nil {
		return false, err
	}
	previousBlock := blockchain.GetCurrentBlock()
	err = newBlock.IsValid(previousBlock)
	if err != nil {
		return false, err
	}
	mutex.Lock()
	blockchain.Chain = append(blockchain.Chain, *newBlock)
	mutex.Unlock()
	return true, nil
}

//GetChain - Returns the current chain as json in blockchain
// TODO - this should return immutable objects
func GetChain() (*string, error) {
	chainBytes, err := json.Marshal(blockchain)
	if err != nil {
		return nil, err
	}
	chainJSON := string(chainBytes)
	return &chainJSON, nil
}

/******* Core functions of the chain **********/

// GenerateBlock - Generates a new block for the given data
// functions critical sections are mutexed
func (bc *BC) GenerateBlock(data interface{}) (*Block, error) {
	currentBlock := bc.GetCurrentBlock()
	mutex.Lock()
	blockIndex := currentBlock.Index + 1
	mutex.Unlock()
	now := time.Now().String()
	newBlock := Block{blockIndex, &now, data, nil, currentBlock.Hash}
	hash, err := newBlock.GetHash()
	if err != nil {
		return nil, err
	}
	newBlock.Hash = hash
	return &newBlock, nil
}

// GetCurrentBlock - Fetches the current block in the chain
// functions critical sections are mutexed
func (bc *BC) GetCurrentBlock() Block {
	mutex.Lock()
	currentBlock := bc.Chain[len(bc.Chain)-1]
	mutex.Unlock()
	return currentBlock
}
