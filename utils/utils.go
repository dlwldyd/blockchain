package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err) // 에러 있으면 exit 1로 프로세스 강제 종료
	}
}

func ToBytes(i interface{}) []byte { // interface{} 는 java 의 Object 처럼 어떤 타입이든 파라미터르 받을 수 있다.
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	HandleErr(encoder.Encode(i))
	return buffer.Bytes()
}
