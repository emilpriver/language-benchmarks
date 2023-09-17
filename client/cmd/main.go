package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

var hosts = []string{
	"http://localhost:3000",
}

type Result struct {
	second   int64
	requests int64
}

/*
* TODO:
* - 1. Send initial request to elton
* - 2. Pull result until job is Done
* - 3. Send result to channel
* - 4. Read result and concat into a bigger slice of results containing all clients result
* - 5. Write to CSV
 */

var ch = make(chan Result)

func main() {
	var wg sync.WaitGroup

	for _, url := range hosts {
		go func() {
			wg.Add(1)

			payload := strings.NewReader(`{
				"method": "GET",
				"threads": 16,
				"connections": 1000,
				"seconds": 10,
				"start_at": "2023-09-17T10:16:34.675Z",
				"url": "https://httpbin.org/ip", 
				"content_type": "application/json",
				"body": ""
			}`)
			// TODO:: change ip
			res, err := http.Post(url, "application/json", payload)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(body))

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("hello")
}
