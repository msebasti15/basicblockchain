package model

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{newGenesisBlock()}}
}

func (b *BlockChain) ReplaceChain(newChain []*Block) {
	if len(newChain) > len(b.Blocks) {
		b.Blocks = newChain
	}
}
