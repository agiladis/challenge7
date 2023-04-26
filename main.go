package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type JsonPayload struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	var jsonPayload = JsonPayload{
		Water: rand.Intn(100) + 1,
		Wind:  rand.Intn(100) + 1,
	}

	jsonByte, err := json.Marshal(jsonPayload)
	if err != nil {
		panic(err)
	}

	var client http.Client

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(jsonByte))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
