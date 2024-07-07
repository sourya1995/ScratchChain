package blockchain

import (
	"crypto/md5"
	"bytes"
	"math/rand"
	"time"
)


type Block struct {
	Hash string
	Data string
	PrevHash string
	Nonce int
}

func (b *Block) ComputeHash() {
	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
	computedHash := md5.Sum(concatenatedData)
	b.Hash = string(computedHash[:])
}

func CreateBlock(data string, prevHash string) *Block {
	rand.Seed(time.Now().UnixNano())
	initialNonce := rand.Intn(1000000)
	block := &Block{"", data, prevHash, initialNonce}
	newProof := NewProofOfWork(block)
	nonce, hash := newProof.MineBlock()
	block.Hash = string(hash[:])
	block.Nonce = nonce
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", "")
}

