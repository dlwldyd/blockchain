package db

import (
	"github.com/boltdb/bolt"
	"github.com/dlwldyd/coin/utils"
)

const (
	// bolt는 key/value DB 이다.
	// bucket 은 RDMS 에서 테이블과 비슷함
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		// parameter 0 : 연결할 데이터베이스 이름, 1 : read write 권한, 2 : 옵션
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(tx *bolt.Tx) error {
			//read, write transaction 을 생상한다.
			_, err = tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)
	}
	return db
}
