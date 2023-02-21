package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/dlwldyd/coin/blockchain"
)

const port string = ":3000";

type homeData struct {
	Title string
	Blocks []*blockchain.Block
}

func main() {
	blockchain := blockchain.GetInstance();
	blockchain.AddBlock("Second Block");
	blockchain.AddBlock("Third Block");

	blockchain.ShowAllBlocks();

	// localhost:3000/ 으로 들어오는 요청에 대한 핸들러
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, "hello world"); // fmt.Print는 console에 출력하지만 fmt.Fprint는 파라미터로 들어가는 writer에 출력한다.
		
		
		// tmpl, err := template.ParseFiles("templates/home.html");
		// if err != nil {
		// 	panic(err)
		// }
		tmpl := template.Must(template.ParseFiles("templates/home.html")) // 위의 주석된 코드 4줄과 같은 기능을 함

		data := homeData{"Home", blockchain.AllBlocks()};
		tmpl.Execute(w, data);
	})

	// 스프링에서 메인메서드라 보면 된다.
	log.Fatal(http.ListenAndServe(port, nil));
}