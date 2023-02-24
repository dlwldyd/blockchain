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

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET" : // "/add" 로 get 요청을 받았을 때
		templates.ExecuteTemplate(w, "add", nil) //add.gohtml을 렌더링
	case "POST" : // "/add" 로 post 요청을 받았을 때
		r.ParseForm() // form 데이터를 가져오기 전에 실행해야한다.
		data := r.Form.Get("blockData") //form 데이터 중 blockData를 가져온다.
		blockchain.GetInstance().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect) // "/"로 리다이렉트
	}
}

func main() {
	blockchain := blockchain.GetInstance();
	// blockchain.AddBlock("Second Block");
	// blockchain.AddBlock("Third Block");

	// blockchain.ShowAllBlocks();

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
		templates.ExecuteTemplate(w, "home", data) // home.gohtml을 렌더링
	})
	http.HandleFunc("/add", add)

	// 스프링에서 메인메서드라 보면 된다.
	log.Fatal(http.ListenAndServe(port, nil));
}