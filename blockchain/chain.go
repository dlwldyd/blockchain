/*
블록체인을 핸들링하기 위한 패키지 이다.
블록체인을 가져오고 블록체인에 블록을 추가할 수 있다.
블록체인은 싱글톤으로 구현되어 하나의 블록체인만 존재한다.
*/
package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/dlwldyd/coin/db"
	"github.com/dlwldyd/coin/utils"
	"sync" // 동기화 처리를 위한 패키지
)

type blockchain struct {
	//blocks []*Block
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var bc *blockchain

var once sync.Once

func (b *blockchain) restore(data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	// 파라미터로 포인터가 들어가야함 왜냐하면 해당 메모리 값에 data를 디코딩(파싱) 하기 때문에
	err := decoder.Decode(b)
	utils.HandleErr(err)
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

/*
블록체인 인스턴스를 가져온다.
*/
func GetInstance() *blockchain {
	if bc == nil {
		once.Do(func() { // 몇개의 쓰레드, goRoutine이 동작하든 해당 함수는 단 한번만 실행된다.
			bc = &blockchain{"", 0}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				bc.AddBlock("Genesis Block")
			} else {
				fmt.Println("....")
				bc.restore(checkpoint)
			}
		})
	}
	fmt.Printf("%d\n", bc.Height)
	return bc
}

/*
블록체인의 마지막 블록의 해시 값을 가져온다.
*/
//func (bc *blockchain) getPrevHash() string {
//	blockLength := len(bc.blocks)
//	if blockLength == 0 {
//		return ""
//	}
//	return bc.blocks[blockLength-1].Hash
//}

/*
데이터와 이전 블록의 해시 값을 합해 해시를 계산한다.
*/
//func getHash(Data, PrevHash string) string {
//	return fmt.Sprintf("%x", sha256.Sum256([]byte(Data+PrevHash)))
//}

/*
블록을 생성한다.
*/
//func createBlock(Data string) *Block {
//	newBlock := Block{Data, "", bc.getPrevHash(), len(GetInstance().blocks) + 1}
//	newBlock.Hash = getHash(newBlock.Data, newBlock.PrevHash)
//	return &newBlock
//}

/*
블록체인에 블록을 추가한다.
*/
//func (bc *blockchain) AddBlock(Data string) {
//	newBlock := createBlock(Data)
//	bc.blocks = append(bc.blocks, newBlock)
//}

/*
블록체인에 있는 모든 블록들을 출력한다.
*/
//func (bc *blockchain) ShowAllBlocks() {
//	for _, Block := range bc.blocks {
//		fmt.Printf("Data : %s\n", Block.Data)
//		fmt.Printf("Hash : %s\n", Block.Hash)
//		fmt.Printf("PrevHash : %s\n", Block.PrevHash)
//		fmt.Println("-------")
//	}
//}

/*
블록체인이 가지고 있는 모든 블록들을 배열로 반환한다.
*/
//func (b *blockchain) AllBlocks() []*Block {
//	return b.blocks
//}

//var ErrNotFound = errors.New("Block not found")
//
//func (b *blockchain) GetBlock(height int) (*Block, error) {
//	if height > len(b.blocks) {
//		return nil, ErrNotFound
//	}
//	return b.blocks[height-1], nil
//}
