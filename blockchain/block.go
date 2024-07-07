package blockchain

import (
	"bytes"
	"crypto/md5"
	"math/rand"
	"time"
)

type Block struct {
	Hash         string
	Data         string
	PrevHash     string
	Nonce        int
	Transactions []*Transaction
}

func (b *Block) ComputeHash() {
	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
	computedHash := md5.Sum(concatenatedData)
	b.Hash = string(computedHash[:])
}

func CreateBlock(data string, prevHash string, transactions []*Transaction) *Block {
	rand.Seed(time.Now().UnixNano())
	initialNonce := rand.Intn(1000000)
	block := &Block{"", data, prevHash, initialNonce, transactions}
	newProof := NewProofOfWork(block)
	nonce, hash := newProof.MineBlock()
	block.Hash = string(hash[:])
	block.Nonce = nonce
	return block
}

func Genesis() *Block {
	coinbaseTransaction := &Transaction{
		Sender:    "Coinbase",
		Recipient: "Genesis",
		Amount:    0.0,
		Coinbase:  true,
	}
	return CreateBlock("Genesis", "", []*Transaction{coinbaseTransaction})
}
