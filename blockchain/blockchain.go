/*
블록체인을 핸들링하기 위한 패키지 이다. 
블록체인을 가져오고 블록체인에 블록을 추가할 수 있다. 
블록체인은 싱글톤으로 구현되어 하나의 블록체인만 존재한다.
*/
package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync" // 동기화 처리를 위한 패키지
)

type block struct {
	data     string;
	hash     string;
	prevHash string;
}

type blockchain struct {
	blocks []*block;
}

var bc *blockchain;
var once sync.Once;

/*
블록체인 인스턴스를 가져온다.
*/
func GetInstance() *blockchain {
	if bc == nil {
		once.Do(func() { // 몇개의 쓰레드, goRoutine이 동작하든 해당 함수는 단 한번만 실행된다.
			bc = &blockchain{};
			bc.AddBlock("Genesis Block");
		})
	}
	return bc;
}

/*
블록체인의 마지막 블록의 해시 값을 가져온다.
*/
func (bc *blockchain) getPrevHash() string {
	blockLength := len(bc.blocks)
	if blockLength == 0 {
		return ""
	}
	return bc.blocks[blockLength-1].hash
}

/*
데이터와 이전 블록의 해시 값을 합해 해시를 계산한다.
*/
func getHash(data, prevHash string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data + prevHash)));
}

/*
블록을 생성한다.
*/
func createBlock(data string) *block {
	newBlock := block{data, "", bc.getPrevHash()}
	newBlock.hash = getHash(newBlock.data, newBlock.prevHash);
	return &newBlock;
}

/*
블록체인에 블록을 추가한다.
*/
func (bc *blockchain) AddBlock(data string) {
	newBlock := createBlock(data);
	bc.blocks = append(bc.blocks, newBlock);
}

/*
블록체인에 있는 모든 블록들을 출력한다.
*/
func (bc *blockchain) ShowAllBlocks() {
	for _, block := range bc.blocks {
		fmt.Printf("data : %s\n", block.data);
		fmt.Printf("hash : %s\n", block.hash);
		fmt.Printf("prevHash : %s\n", block.prevHash);
		fmt.Println("-------");
	}
}