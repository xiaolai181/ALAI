package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

func main() {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(s))
}
