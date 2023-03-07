package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dlwldyd/coin/blockchain"
	"github.com/dlwldyd/coin/utils"
	"github.com/gorilla/mux"
)

var port string

type url string // URL이라는 타입을 만듦

// Marshal이나 UnMarshal을 할 때 어떻게 변환되는지를 정의하는 메서드이다.
// 반환타입이 ([]byte, error)이고, 파라미터가 없어야 한다.
func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	// 구조체의 앞글자가 소문자로 시작하면 해당 필드는 export하지 않기 때문에 무조건 대문자로 시작해야한다. 
	// 만약 json 필드를 전부 소문자로(혹은 다른 글자로) 하고 싶으면 아래와 같이 백틱을 사용해서 json에서 어떻게 보여질 지 명시할 수 있다.
	// Field Tags를 검색하면 백틱에 들어가는 것들에 대한 정보를 더 얻을 수 있다.
	URL url `json:"url"`
	Method string `json:"method"`
	Description string `json:"description"`

	Payload string `json:"payload,omitempty"` // omitempty가 들어가면 만약 값이 비어있다면 아예 json에 값을 넣지 않는다.
}

type addBlockBody struct {
	Message string
}

// fmt로 출력할 때 호출한다. 자바의 Object 객체의 toString 메서드라 보면된다.
// fmt를 통해 출력하려는 객체의 메서드로 String()이 존재하면(반환타입이 string이고 파라미터가 없어야함) 자동으로 출력한다.
func (u urlDescription) String() string {
	return "hello i'am url description"
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription {
		{
			URL: url("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: url("/blocks"),
			Method: "GET",
			Description: "See All Blocks",
		},
		{
			URL: url("/blocks"),
			Method: "POST",
			Description: "Add A Block",
			Payload: "data:string",
		},
		{
			URL: url("/blocks/{height}"),
			Method: "GET",
			Description: "See A Block",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	
	// b, err := json.Marshal(data) // 구조체를 json으로 바꿔줌(byte array로 바뀜)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s\n", b)
	json.NewEncoder(rw).Encode(data) // 위의 3줄과 같은 기능을 한다.
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	switch r.Method {
	case "GET" :
		json.NewEncoder(rw).Encode(blockchain.GetInstance().AllBlocks())
	case "POST" : 
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetInstance().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block := blockchain.GetInstance().GetBlock(id)
	json.NewEncoder(rw).Encode(block)
}

func Start(aPort int) {
	// 같은 url, 다른 포트를 사용하고 싶으면 DefaultServeMux를 사용할 수 없다.
	// Mux는 HandleFunc 함수에서 url과 handler를 매핑시켜주는 역할을 한다.
	// 만약 사용할 Mux를 ListenAndServe 함수에 파라미터로 넣지 않는다면 DefaultServeMux를 사용한다.
	// handler := http.NewServeMux()
	
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	
	// http.HandleFunc("/", documentation)
	// http.HandleFunc("/blocks", blocks)
	// handler.HandleFunc("/", documentation)
	// handler.HandleFunc("/blocks", blocks)
	
	// gorilla mux는 핸들러가 어떤 http method를 다루는지 지정할 수 있다.
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")
	fmt.Printf("Listening on http://localhost%s\n", port)

	// http 패키지에 있는 HandleFunc 함수가 아니라(DefaultServeMux 사용) 다른 Mux를 사용한다면 2번째 파라미터로 handler를 넣어줘야한다.
	log.Fatal(http.ListenAndServe(port, router))
}