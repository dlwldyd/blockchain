package main

import (
	"github.com/dlwldyd/coin/explorer"
	"github.com/dlwldyd/coin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}