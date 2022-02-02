## tpool

Golang Implementation of Worker Pool/ Thread Pool

#

## What is Worker pool/ Thread Pool
From Rust's docs https://doc.rust-lang.org/book/ch20-02-multithreaded.html

A thread pool is a group of spawned threads that are waiting and ready to handle a task. When the program receives a new task, it assigns one of the threads in the pool to the task, and that thread will process the task. The remaining threads in the pool are available to handle any other tasks that come in while the first thread is processing. When the first thread is done processing its task, it’s returned to the pool of idle threads, ready to handle a new task. A thread pool allows you to process connections concurrently, increasing the throughput of your server.

We’ll limit the number of threads in the pool to a small number to protect us from Denial of Service (DoS) attacks; if we had our program create a new thread for each request as it came in, someone making 10 million requests to our server could create havoc by using up all our server’s resources and grinding the processing of requests to a halt.

Rather than spawning unlimited threads, we’ll have a fixed number of threads waiting in the pool. As requests come in, they’ll be sent to the pool for processing. The pool will maintain a queue of incoming requests. Each of the threads in the pool will pop off a request from this queue, handle the request, and then ask the queue for another request. With this design, we can process N requests concurrently, where N is the number of threads. If each thread is responding to a long-running request, subsequent requests can still back up in the queue, but we’ve increased the number of long-running requests we can handle before reaching that point.

### Usage

Simulate heavy job that takes 1 second each to complete the job. With 3 worker it only takes 2 seconds instead 4 second

try to modify
```
threadPool := tpool.NewThreadPool(3)
```
to

```
threadPool := tpool.NewThreadPool(4)
```

and see what happen

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/telkomdev/tpool"
)

type Arg struct {
	X uint
	Y uint
}

func main() {

	job1 := tpool.NewJob(Arg{
		X: 5, Y: 10,
	}, func(arg tpool.JobArg, res chan<- tpool.Result) error {
		a := arg.(Arg)
		r := a.X * a.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job2 := tpool.NewJob(Arg{
		X: 25, Y: 4,
	}, func(arg tpool.JobArg, res chan<- tpool.Result) error {
		a := arg.(Arg)
		r := a.X * a.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job3 := tpool.NewJob(Arg{
		X: 100, Y: 4,
	}, func(arg tpool.JobArg, res chan<- tpool.Result) error {
		a := arg.(Arg)
		r := a.X * a.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job4 := tpool.NewJob(Arg{
		X: 5, Y: 5,
	}, func(arg tpool.JobArg, res chan<- tpool.Result) error {
		a := arg.(Arg)
		r := a.X * a.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	// start
	start := time.Now()

	jobs := []tpool.Job{job1, job2, job3, job4}

	ctx, cancel := context.WithCancel(context.Background())
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

```

### Full Example
Open the `_example` folder