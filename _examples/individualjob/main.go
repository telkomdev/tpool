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
	}, func(arg Arg, res chan<- uint) error {
		r := arg.X * arg.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job2 := tpool.NewJob(Arg{
		X: 25, Y: 4,
	}, func(arg Arg, res chan<- uint) error {
		r := arg.X * arg.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job3 := tpool.NewJob(Arg{
		X: 100, Y: 4,
	}, func(arg Arg, res chan<- uint) error {
		r := arg.X * arg.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job4 := tpool.NewJob(Arg{
		X: 5, Y: 5,
	}, func(arg Arg, res chan<- uint) error {
		r := arg.X * arg.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job5 := tpool.NewJob(Arg{
		X: 100, Y: 100,
	}, func(arg Arg, res chan<- uint) error {
		r := arg.X * arg.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job6 := tpool.NewJob(Arg{
		X: 2, Y: 2,
	}, func(arg Arg, res chan<- uint) error {
		r := arg.X * arg.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job7 := tpool.NewJob(Arg{
		X: 25, Y: 2,
	}, func(arg Arg, res chan<- uint) error {
		r := arg.X * arg.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job8 := tpool.NewJob(Arg{
		X: 10, Y: 2,
	}, func(arg Arg, res chan<- uint) error {
		r := arg.X * arg.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	// start
	start := time.Now()

	jobs := []tpool.Job[Arg, uint]{job1, job2, job3, job4, job5, job6, job7, job8}

	ctx, cancel := context.WithCancel(context.Background())
	// ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*800)
	defer func() { cancel() }()

	threadPool := tpool.NewThreadPool[Arg, uint](2)

	// send individual job
	go func() {
		for _, j := range jobs {
			threadPool.Jobs() <- j
		}
		close(threadPool.Jobs())
	}()

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
