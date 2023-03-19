package main

import "github.com/dlwldyd/coin/blockchain"

func main() {
	blockchain.GetInstance().AddBlock("first")
	blockchain.GetInstance().AddBlock("second")
	blockchain.GetInstance().AddBlock("third")
}
