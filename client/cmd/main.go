package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

var hosts = []string{
	"http://localhost:3000",
}

type Result struct {
	ID         string `json:"id"`
	TestID     string `json:"test_id"`
	Second     int    `json:"second"`
	Requests   int    `json:"requests"`
	ErrorCodes string `json:"error_codes"`
}

type Response struct {
	ID          string   `json:"id"`
	URL         string   `json:"url"`
	Method      string   `json:"method"`
	ContentType string   `json:"content_type"`
	Status      string   `json:"status"`
	Body        string   `json:"body"`
	CreatedAt   string   `json:"created_at"`
	FinishedAt  string   `json:"finished_at"`
	Results     []Result `json:"results"`
}

/*
* TODO:
* - 1. Send initial request to elton
* - 2. Pull result until job is Done
* - 3. Send result to channel
* - 4. Read result and concat into a bigger slice of results containing all clients result
* - 5. Write to CSV
 */

var ch = make(chan []Result)

func main() {
	var wg sync.WaitGroup

	for _, url := range hosts {
		wg.Add(1)

		go func(remoteUrl string) {
			defer wg.Done()

			payload := strings.NewReader(`{
				"method": "GET",
				"tasks": 16,
				"seconds": 3,
				"start_at": "2023-09-17T10:16:34.675Z",
				"url": "https://httpbin.org/ip", 
				"content_type": "application/json",
				"body": ""
			}`)
			// TODO:: change ip
			res, err := http.Post(remoteUrl, "application/json", payload)
			if err != nil {
				fmt.Println("sending post request")
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println("reading post body")
				fmt.Println(err)
				return
			}

			fmt.Println(string(body))

			var response Response
			err = json.Unmarshal(body, &response)
			if err != nil {
				fmt.Println("Unmarshal post request")
				fmt.Println(err)
				return
			}

			for {
				getTest, err := http.Get(fmt.Sprintf("%s/%s", remoteUrl, response.ID))
				if err != nil {
					fmt.Println("Sending get request")
					fmt.Println(err)
					return
				}

				body, err := io.ReadAll(getTest.Body)
				if err != nil {
					fmt.Println("reading get request")
					fmt.Println(err)
					return
				}
				fmt.Println(string(body))

				var getTestResponse Response
				err = json.Unmarshal(body, &getTestResponse)
				if err != nil {
					fmt.Println("unmarshal get request")
					return
				}

				if getTestResponse.Status == "PROCESSING" {
					time.Sleep(1 * time.Second)
					continue
				}

				if getTestResponse.Status == "FINISHED" {
					fmt.Println("hello")
					ch <- getTestResponse.Results

					break
				}

				fmt.Println(fmt.Sprintf("unexpected status :%s", getTestResponse.Status))
				break
			}

			fmt.Println("break the loop")
		}(url)
	}

	wg.Wait()

	for r := range ch {
		fmt.Println("hello")
		fmt.Println(r)
	}

	fmt.Println("hello")
}
