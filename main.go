package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	apiUrl = "https://api.codex.jaagrav.in"
)

type Response struct {
	TimeStamp int    `json:"timeStamp"`
	Status    int    `json:"status"`
	Output    string `json:"output"`
	Error     string `json:"error"`
	Language  string `json:"language"`
	Info      string `json:"info"`
}

func read(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func main() {
	var (
		FilePath  = flag.String("f", "", "input file")
		Extension = flag.String("e", "", "code programming language")
		Input     = flag.String("i", "", "code input")
	)

	flag.Parse()

	content, err := read(*FilePath)
	if err != nil {
		panic(err)
	}

	queryParams := url.Values{
		"code":     {content},
		"language": {*Extension},
		"input":    {*Input},
	}

	payload := bytes.NewReader([]byte(queryParams.Encode()))
	client := http.Client{}

	req, _ := http.NewRequest(http.MethodPost, apiUrl, payload)
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	var r Response

	bytesString, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(bytesString, &r); err != nil {
		panic(err)
	}

	log.Println(r)
}
