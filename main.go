package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type JsonPayload struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {

	var URL string = "https://jsonplaceholder.typicode.com/posts"
	var METHOD string = "POST"

	after := time.Now().Add(60 * time.Second)

	for {
		now := time.Now()

		var jsonPayload = JsonPayload{
			Water: rand.Intn(100) + 1,
			Wind:  rand.Intn(100) + 1,
		}

		jsonResponse, err := DoRequest(METHOD, URL, jsonPayload)
		if err != nil {
			panic(err)
		}

		// JSON response
		fmt.Println(jsonResponse)
		// Water status
		if jsonPayload.Water < 5 {
			fmt.Println("status water : aman")
		} else if jsonPayload.Water > 8 {
			fmt.Println("status water : bahaya")
		} else {
			fmt.Println("status water : siaga")
		}
		// Wind status
		if jsonPayload.Water < 6 {
			fmt.Println("status water : aman")
		} else if jsonPayload.Water > 15 {
			fmt.Println("status water : bahaya")
		} else {
			fmt.Println("status water : siaga")
		}
		// Pagination
		fmt.Println("")
		fmt.Println("===================================")
		fmt.Println("")

		time.Sleep(15 * time.Second)

		if now.After(after) {
			break
		}
	}
}

func DoRequest(method, url string, payload any) (string, error) {
	var client http.Client

	jsonByte, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonByte))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
