package tester

import (
	"log"
	"net/http"
	"time"
)

func StressTest(cfg *Configs) Report {
	requestChannel := make(chan string, cfg.Requests)
	statusChannel := make(chan int)

	for range cfg.Concurrency {
		go makeRequest(requestChannel, statusChannel)
	}

	startTime := time.Now()
	for range cfg.Requests {
		requestChannel <- cfg.Url
	}

	var reportResult Report
	for status := range statusChannel {
		reportResult.TotalRequests++

		if status != 200 {
			reportResult.TotalAnotherStatus++
		} else {
			reportResult.TotalStatusOk++
		}

		if reportResult.TotalRequests == cfg.Requests {
			close(statusChannel)
		}
	}
	reportResult.TotalTime = time.Since(startTime)

	return reportResult
}

func makeRequest(requestChannel chan string, statusChannel chan int) {
	client := &http.Client{Timeout: 0}

	for url := range requestChannel {
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		request.Header.Set("User-Agent", "stress-test/1.0")

		response, err := client.Do(request)
		if err != nil {
			log.Fatal("erro de request", err)
		}

		statusChannel <- response.StatusCode
	}
}
