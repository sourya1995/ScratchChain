package main

import (
	"fmt"
	"strconv"
	"blockchain/blockchain"
)

func main() {
    chain := blockchain.InitBlockChain()

    chain.AddBlock("Block 1", "Alice", []*blockchain.Transaction{
        {Sender: "Alice", Recipient: "Bob", Amount: 1.5},
        {Sender: "Alice", Recipient: "Charlie", Amount: 19.5},
    })

    chain.AddBlock("Block 2", "Bob", []*blockchain.Transaction{
        {Sender: "Bob", Recipient: "Charlie", Amount: 2.3},
    })

    for _, block := range chain.Blocks {
        fmt.Printf("Previous hash: %x\n", block.PrevHash)
        fmt.Printf("Data in Block: %s\n", block.Data)
        fmt.Printf("Hash of block: %x\n", block.Hash)

        pow := blockchain.NewProofOfWork(block)
        fmt.Printf("IsValidPoW: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()

        fmt.Println("Transactions:")

        for _, tx := range block.Transactions {
            fmt.Printf("Sender: %s\n", tx.Sender)
            fmt.Printf("Receiver: %s\n", tx.Recipient)
            fmt.Printf("Amount: %f\n", tx.Amount)
            fmt.Printf("Coinbase: %t\n", tx.Coinbase)
            fmt.Println()
        }
        fmt.Println()
    }
}
