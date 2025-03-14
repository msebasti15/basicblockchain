package main

import (
	"basicblockchain/internal/model"
	"fmt"
)

func main() {
	bc := model.NewBlockChain()

	bc.AddBlock("Send 1 BTC to Me")
	bc.AddBlock("Send 2 more BTC to Me")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		fmt.Println()
	}
}
