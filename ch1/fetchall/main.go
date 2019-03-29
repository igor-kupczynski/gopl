package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()

	results := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, results)
	}

	for range os.Args[1:] {
		fmt.Println(<-results)
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs Total\n", secs)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("while fetching [%s]: %v", url, err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while readin [%s]: %v", url, err)
		return
	}
	_ = resp.Body.Close()

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
