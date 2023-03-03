package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
const port string = ":3000"

type URL string // URL이라는 타입을 만듦

// Marshal이나 UnMarshal을 할 때 어떻게 변환되는지를 정의하는 메서드이다.
// 반환타입이 ([]byte, error)이고, 파라미터가 없어야 한다.
func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	// 구조체의 앞글자가 소문자로 시작하면 해당 필드는 export하지 않기 때문에 무조건 대문자로 시작해야한다. 
	// 만약 json 필드를 전부 소문자로(혹은 다른 글자로) 하고 싶으면 아래와 같이 백틱을 사용해서 json에서 어떻게 보여질 지 명시할 수 있다.
	// Field Tags를 검색하면 백틱에 들어가는 것들에 대한 정보를 더 얻을 수 있다.
	URL URL `json:"url"`
	Method string `json:"method"`
	Description string `json:"description"`

	Payload string `json:"payload,omitempty"` // omitempty가 들어가면 만약 값이 비어있다면 아예 json에 값을 넣지 않는다.
}

// fmt로 출력할 때 호출한다. 자바의 Object 객체의 toString 메서드라 보면된다.
// fmt를 통해 출력하려는 객체의 메서드로 String()이 존재하면(반환타입이 string이고 파라미터가 없어야함) 자동으로 출력한다.
func (u URLDescription) String() string {
	return "hello i'am url description"
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription {
		{
			URL: URL("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: URL("/blocks"),
			Method: "POST",
			Description: "Add A Block",
			Payload: "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	
	// b, err := json.Marshal(data) // 구조체를 json으로 바꿔줌(byte array로 바뀜)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s\n", b)
	json.NewEncoder(rw).Encode(data) // 위의 3줄과 같은 기능을 한다.
}

func main() {
	// explorer.Start()
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}