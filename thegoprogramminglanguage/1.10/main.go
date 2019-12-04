package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

var reg *regexp.Regexp

func main() {
	reg, _ = regexp.Compile("[^a-zA-Z0-9]+")

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	file, err := os.Create(reg.ReplaceAllString(url, "") + strconv.FormatInt(time.Now().Unix(), 10) + ".txt")
	if err != nil {
		ch <- fmt.Sprintf("While creating file for %s: %v", url, err)
		return
	}
	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
