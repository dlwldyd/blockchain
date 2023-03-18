package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"github.com/dlwldyd/coin/db"
	"github.com/dlwldyd/coin/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func createBlock(data string, prevHash string, height int) *Block {
	block := Block{
		data, "", prevHash, height,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%s", sha256.Sum256([]byte(payload)))
	block.persist()
	return &block
}

func (b *Block) toBytes() []byte {
	var blockBuffer bytes.Buffer
	encoder := gob.NewEncoder(&blockBuffer)
	err := encoder.Encode(b)
	utils.HandleErr(err)
	return blockBuffer.Bytes()
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, b.toBytes())
}
