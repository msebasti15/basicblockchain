package model

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Index         int64
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) CalculateHash() []byte {
	index := []byte(strconv.FormatInt(b.Index, 10))
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{index, timestamp, b.PrevBlockHash, b.Data}, []byte{})
	hash := sha256.Sum256(headers)

	return hash[:]
}

func (b *Block) SetHash() {
	b.Hash = b.CalculateHash()
}

func newGenesisBlock() *Block {
	block := &Block{Index: 0, Timestamp: time.Now().Unix(), Data: []byte("Genesis Block"), PrevBlockHash: nil, Hash: []byte{}}
	block.SetHash()

	return block
}

func NewBlock(data string, prevBlock *Block) *Block {
	block := &Block{Index: prevBlock.Index + 1, Timestamp: time.Now().Unix(), Data: []byte(data), PrevBlockHash: prevBlock.Hash, Hash: []byte{}}
	block.SetHash()

	return block
}

func IsBlockValid(newBlock, prevBloc *Block) bool {
	if prevBloc.Index+1 != newBlock.Index {
		return false
	}

	if !bytes.Equal(prevBloc.Hash, newBlock.PrevBlockHash) {
		return false
	}

	if !bytes.Equal(newBlock.Hash, newBlock.CalculateHash()) {
		return false
	}

	return true
}
