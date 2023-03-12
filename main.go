package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("welcome !!\n\n")
	fmt.Printf("please use following commands:\n\n")
	fmt.Printf("explorer:	Start the HTML Explorer\n")
	fmt.Printf("rest:		Start the REST API(recommended\n\n")
	os.Exit(0) // 프로그램 종, 종료시 아무 메세지 없음
	//os.Exit(1) // 프로그램 종료, exit status 1 이라는 메세지 출력하면서 종료
}

func main() {
	// go explorer.Start(3000)
	// rest.Start(4000)
	if len(os.Args) < 2 { // os.Args는 Array of string 인데 0번 인덱스에는 프로그램의 이름이 들어가고 그 다음부터는 옵션 및 커맨드가 들어간다
		usage()
	}

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		fmt.Println("Start REST API")
	default:
		usage()
	}
}
