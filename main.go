package main

// reference: https://medium.com/@zach_4342/dependency-injection-in-golang-e587c69478a8#.3eu1gzfam

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type HttpClient interface {
	Get(string) (*http.Response, error)
}

var (
	url string
)

func init() {
	isTest := strings.HasSuffix(os.Args[0], ".test")
	if !isTest {
		flag.StringVar(&url, "url", "http://google.com", "Of which URL should we print the response?")
		flag.Parse()
	}
}

func main() {
	client := &http.Client{}
	err := send(client, url)

	if err != nil {
		panic(err)
	}
}

func send(client HttpClient, link string) error {
	response, err := client.Get(link)

	if err != nil {
		return err
	}

	if response == nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
