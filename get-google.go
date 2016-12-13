package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("https://google.com/")
	if err != nil {
		fmt.Errorf("Problem requesting google, %s", err)
	}

	print(resp)
	print("\n")
	fmt.Print(resp)
	print("\n")
	print("\n")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Errorf("Problem reading body, %s", err)
	}

	fmt.Printf("%s", body)
}
