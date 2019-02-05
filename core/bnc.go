package core

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
	"reflect"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

var (
	// Blockchain - Simple chain of blocks
	Blockchain []Block

	genesisHash     = "dd102105cdd3a38ce77562d856cfce6d"
	genesisSeedData = map[string]interface{}{"whoami": 42}
)

// Block - Basic datastructure of a simple block
type Block struct {
	Index             int         `json:"index"`
	Timestamp         *string     `json:"timestamp"`
	Data              interface{} `json:"data"`
	Hash              *string     `json:"hash"`
	PreviousBlockHash *string     `json:"previous_block_hash"`
}

func hashOfData(data interface{}) string {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Ptr {
		if !v.CanAddr() {
			return ""
		}

		v = v.Addr()
	}
	size := unsafe.Sizeof(v.Interface())
	b := (*[1 << 10]uint8)(unsafe.Pointer(v.Pointer()))[:size:size]

	h := md5.New()
	return base64.StdEncoding.EncodeToString(h.Sum(b))
}

// Hash - Caclulates the hash of block, will work for genesis block as well
func Hash(b *Block) string {
	digest := strconv.Itoa(b.Index) + *b.Timestamp + hashOfData(b.Data)
	//Doble check this is indeed the genesis block irrespective of Valid block
	//PreviousBlockHash can be nil only in genesis block
	if b.PreviousBlockHash == nil && *b.Hash != genesisHash && b.Data != nil {
		//Fatal, PreviousHash can be nil only for genesis block
		log.Fatal(errors.New("Non genesis block should always have PreviousBlockHash"))
	} else {
		digest = digest + *b.PreviousBlockHash
	}
	hash := sha256.New()
	hash.Write([]byte(digest))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

var mutex = &sync.Mutex{}

func getGenesisBlock() Block {
	now := time.Now()
	genTimeStamp := now.String()
	block := Block{0, &genTimeStamp, nil, &genesisHash, nil}
	return block
}

func init() {
	genBlock := getGenesisBlock()
	Blockchain = append(Blockchain, genBlock)
}
