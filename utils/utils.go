package utils

import "log"

func HandleErr(err error) {
	if err != nil {
		log.Panic(err) // 에러 있으면 exit 1로 프로세스 강제 종료
	}
}