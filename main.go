package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/dlwldyd/coin/blockchain"
)

const(
	port string = ":3000";
	templateDir string = "templates/"
)
var templates *template.Template

type homeData struct {
	Title string
	Blocks []*blockchain.Block
}

func main() {
	blockchain := blockchain.GetInstance();
	blockchain.AddBlock("Second Block");
	blockchain.AddBlock("Third Block");

	blockchain.ShowAllBlocks();

	// parseGlob는 여러개의 템플릿을 한번에 로드할 때 사용한다.
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	

	// localhost:3000/ 으로 들어오는 요청에 대한 핸들러
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, "hello world"); // fmt.Print는 console에 출력하지만 fmt.Fprint는 파라미터로 들어가는 writer에 출력한다.
		
		
		// tmpl, err := template.ParseFiles("templates/home.html");
		// if err != nil {
		// 	panic(err)
		// }
		// tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml")) // 위의 주석된 코드 4줄과 같은 기능을 함
		
		// tmpl.Execute(w, data);
		data := homeData{"Home", blockchain.AllBlocks()};
		templates.ExecuteTemplate(w, "home", data)
	})

	// 스프링에서 메인메서드라 보면 된다.
	log.Fatal(http.ListenAndServe(port, nil));
}