package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Text string `json:"think"`
}

func main() {
	var args = os.Args[1:]
	var fulltext = strings.Join(args, " ")

	postJSON(makeThink(fulltext))
}

func postJSON(message []byte) {
	url := "http://devil:5000/api/think"

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(message))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if error != nil {
		panic(error)
	}

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

}

func makeThink(msg string) []byte {
	message := &Message{Text: msg}

	b, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}

	return b
}
