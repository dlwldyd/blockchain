package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dlwldyd/coin/utils"
)
const port string = ":3000"

type URLDescription struct {
	URL string
	Method string
	Description string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription {
		{
			URL: "/",
			Method: "GET",
			Description: "See Documentation",
		},
	}
	b, err := json.Marshal(data) // 구조체를 json으로 바꿔줌(byte array로 바뀜)
	utils.HandleErr(err)
	fmt.Printf("%s\n", b)
}

func main() {
	// explorer.Start()
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}