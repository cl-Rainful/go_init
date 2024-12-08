package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func httpDemo() {

	client := http.Client{

		Timeout: 1 * time.Second,
	}

	if res, err := client.Get("http://127.0.0.1:5678"); err != nil {
		defer res.Body.Close()
		fmt.Println(res.StatusCode)
		if bs, err := io.ReadAll(res.Body); err == nil {
			fmt.Println(string(bs))
		}
	} else {
		fmt.Println(err)
	}
}

func httpWelcom(ch chan<- string) {

	time.Sleep(2 * time.Second)
	ch <- "welcome"
}

func httpWelcomHandler(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	blocker := make(chan string, 1)
	go httpWelcom(blocker)
	select {
	case v := <-blocker:
		fmt.Fprint(w, v)
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server: ", err)
	}
}
