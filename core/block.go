package core

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"strconv"

	"github.com/heckdevice/gobc/utils"
)

// Block - Basic datastructure of a simple block
type Block struct {
	Index             int         `json:"index"`
	Timestamp         *string     `json:"timestamp"`
	Data              interface{} `json:"data"`
	Hash              *string     `json:"hash"`
	PreviousBlockHash *string     `json:"previous_block_hash"`
}

/******* Core functions of the block **********/

// GetHash - Caclulates and returns the hash of the block
// Data, Index, PreviousBlockHash and Timestamp are mandatory for hashing
func (b *Block) GetHash() (*string, error) {
	// PreviousBlockHash cannot be nil except genesis block
	// Data cannot be nil
	// Index cannot be of genesis block
	// Timestamp cannot be nil
	if b.Index == 0 || b.PreviousBlockHash == nil || b.Timestamp == nil || b.Data == nil {
		err := errors.New("Block invalid for hashing")
		log.Fatal(err)
		return nil, err
	}
	digest := strconv.Itoa(b.Index) + *b.Timestamp + utils.Hasher(b.Data) + *b.PreviousBlockHash
	hash := sha256.New()
	hash.Write([]byte(digest))
	hashed := hash.Sum(nil)
	strHash := hex.EncodeToString(hashed)
	return &strHash, nil
}
