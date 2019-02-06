package core

import (
	"reflect"
	"sync"
	"time"

	"github.com/heckdevice/gobc/utils"
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
	Count       int     `json:"count"`
	Size        uintptr `json:"size"`
	CurrentHash string  `json:"current_hash"`
	Chain       []Block `json:"chain"`
}

func getGenesisBlock() Block {
	now := time.Now()
	genTimeStamp := now.String()
	block := Block{0, &genTimeStamp, nil, &genesisHash, nil}
	return block
}

// Init - Initialize a Blockchain with genesis block
func Init() {
	mutex.Lock()
	if blockchain.Chain == nil || len(blockchain.Chain) == 0 {
		blockchain.Chain = make([]Block, 1, 1)
		genBlock := getGenesisBlock()
		blockchain.Chain[0] = genBlock
		blockchain.Count = len(blockchain.Chain)
		blockchain.Size = uintptr(len(blockchain.Chain)) * reflect.TypeOf(blockchain.Chain).Elem().Size()
		blockchain.CurrentHash = *genBlock.Hash
	}
	mutex.Unlock()
}

// Add - adds new data to the blockchain
// functions critical sections are mutexed
func Add(data interface{}) (bool, *Block, error) {
	newBlock, err := blockchain.GenerateBlock(data)
	if err != nil {
		return false, nil, err
	}
	previousBlock := blockchain.GetCurrentBlock()
	err = newBlock.IsValid(previousBlock)
	if err != nil {
		return false, nil, err
	}
	mutex.Lock()
	blockchain.Chain = append(blockchain.Chain, *newBlock)
	blockchain.Count = len(blockchain.Chain)
	blockchain.Size = uintptr(len(blockchain.Chain)) * reflect.TypeOf(blockchain.Chain).Elem().Size()
	blockchain.CurrentHash = *newBlock.Hash
	mutex.Unlock()
	return true, newBlock, nil
}

//GetChain - Returns the current blockchain as json
// TODO - this should return immutable objects and not json
func GetChain() (*string, error) {
	return utils.InterfaceToJSONString(blockchain)
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
