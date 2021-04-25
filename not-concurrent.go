package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"sync"
	"time"
)

func accessAndTime(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	client := http.Client{}
	resp, _ := client.Get(url)
	defer resp.Body.Close()

	ioutil.ReadAll(resp.Body)

	duration := time.Since(start)

	fmt.Printf("Website is loaded\n", url, duration)
}

func main() {
	sites := []string {
		"https://www.bbc.co.uk",
		"https://www.edition.cnn.com",
		"https://www.theguardian.com/uk",
		"https://old.reddit.com",
	}

	var wg sync.WaitGroup

	for _, value :=	range sites {
		wg.Add(1)
		go accessAndTime(value, &wg)
	}

	wg.Wait()
}