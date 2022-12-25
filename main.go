package main

import "github.com/dlwldyd/coin/blockchain"

func main() {
	blockchain := blockchain.GetInstance();
	blockchain.AddBlock("Second Block");
	blockchain.AddBlock("Third Block");

	blockchain.ShowAllBlocks();
}