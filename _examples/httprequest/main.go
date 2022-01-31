package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/telkomdev/tpool"
)

func httpGet(url string) (*http.Response, error) {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		IdleConnTimeout:     10 * time.Second,
	}

	httpClient := &http.Client{
		//Timeout:   time.Second * 10,
		Transport: transport,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func main() {
	arg := "https://www.w3schools.com/js/default.asp"
	job1 := tpool.NewJob(arg, func(arg tpool.JobArg, res chan<- tpool.Result) error {
		a := arg.(string)
		start := time.Now()
		// todo
		resp, err := httpGet(a)
		if err != nil {
			r := fmt.Sprintf("err = %s", err)
			res <- r
		} else {
			fmt.Printf("job run in %s ms\n", time.Since(start))
			r := fmt.Sprintf("resp = %s", resp.Status)
			res <- r
		}

		return nil
	})

	job2 := tpool.NewJob(arg, func(arg tpool.JobArg, res chan<- tpool.Result) error {
		a := arg.(string)
		start := time.Now()
		// todo
		resp, err := httpGet(a)
		if err != nil {
			r := fmt.Sprintf("err = %s", err)
			res <- r
		} else {
			fmt.Printf("job run in %s ms\n", time.Since(start))
			r := fmt.Sprintf("resp = %s", resp.Status)
			res <- r
		}

		return nil
	})

	job3 := tpool.NewJob(arg, func(arg tpool.JobArg, res chan<- tpool.Result) error {
		a := arg.(string)
		start := time.Now()
		// todo
		resp, err := httpGet(a)
		if err != nil {
			r := fmt.Sprintf("err = %s", err)
			res <- r
		} else {
			fmt.Printf("job run in %s ms\n", time.Since(start))
			r := fmt.Sprintf("resp = %s", resp.Status)
			res <- r
		}

		return nil
	})

	job4 := tpool.NewJob(arg, func(arg tpool.JobArg, res chan<- tpool.Result) error {
		a := arg.(string)
		start := time.Now()
		// todo
		resp, err := httpGet(a)
		if err != nil {
			r := fmt.Sprintf("err = %s", err)
			res <- r
		} else {
			fmt.Printf("job run in %s ms\n", time.Since(start))
			r := fmt.Sprintf("resp = %s", resp.Status)
			res <- r
		}

		return nil
	})

	job5 := tpool.NewJob(arg, func(arg tpool.JobArg, res chan<- tpool.Result) error {
		a := arg.(string)
		start := time.Now()
		// todo
		resp, err := httpGet(a)
		if err != nil {
			r := fmt.Sprintf("err = %s", err)
			res <- r
		} else {
			fmt.Printf("job run in %s ms\n", time.Since(start))
			r := fmt.Sprintf("resp = %s", resp.Status)
			res <- r
		}

		return nil
	})

	job6 := tpool.NewJob(arg, func(arg tpool.JobArg, res chan<- tpool.Result) error {
		a := arg.(string)
		start := time.Now()
		// todo
		resp, err := httpGet(a)
		if err != nil {
			r := fmt.Sprintf("err = %s", err)
			res <- r
		} else {
			fmt.Printf("job run in %s ms\n", time.Since(start))
			r := fmt.Sprintf("resp = %s", resp.Status)
			res <- r
		}

		return nil
	})

	job7 := tpool.NewJob(arg, func(arg tpool.JobArg, res chan<- tpool.Result) error {
		a := arg.(string)
		start := time.Now()
		// todo
		resp, err := httpGet(a)
		if err != nil {
			r := fmt.Sprintf("err = %s", err)
			res <- r
		} else {
			fmt.Printf("job run in %s ms\n", time.Since(start))
			r := fmt.Sprintf("resp = %s", resp.Status)
			res <- r
		}

		return nil
	})

	job8 := tpool.NewJob(arg, func(arg tpool.JobArg, res chan<- tpool.Result) error {
		a := arg.(string)
		start := time.Now()
		// todo
		resp, err := httpGet(a)
		if err != nil {
			r := fmt.Sprintf("err = %s", err)
			res <- r
		} else {
			fmt.Printf("job run in %s ms\n", time.Since(start))
			r := fmt.Sprintf("resp = %s", resp.Status)
			res <- r
		}

		return nil
	})

	// start
	start := time.Now()

	jobs := []tpool.Job{job1, job2, job3, job4, job5, job6, job7, job8}

	ctx, cancel := context.WithCancel(context.Background())
	// ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*800)
	defer func() { cancel() }()

	threadPool := tpool.NewThreadPool(3)
	threadPool.GenerateJobFrom(jobs)

	threadPool.Run(ctx)

	for {
		select {
		case r, ok := <-threadPool.Result():
			if !ok {
				break
			}

			fmt.Println(r)
		case <-threadPool.Done:
			fmt.Println("worker done")
			elapsed := time.Since(start)

			fmt.Printf("all worker run in %s\n", elapsed)
			return
		}
	}

}
