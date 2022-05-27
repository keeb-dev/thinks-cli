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
	// todo: replace with evans thing https://discord.com/channels/976812245843583006/976812246376280087/979554599318982718
	var args = os.Args[1:]
	var fulltext = strings.Join(args, " ")

	postJSON(makeThink(fulltext))
}

func postJSON(message []byte) {
	// todo: need to externalize this
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
	// todo: consider stripping some of the text here and ensuring it is not blank
	message := &Message{Text: msg}

	b, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}

	return b
}
