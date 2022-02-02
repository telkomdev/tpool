package tpool

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func generateJobs() []Job {
	type Arg struct {
		X uint
		Y uint
	}

	job1 := NewJob(Arg{
		X: 5, Y: 10,
	}, func(arg JobArg, res chan<- Result) error {
		a := arg.(Arg)
		r := a.X * a.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job2 := NewJob(Arg{
		X: 25, Y: 4,
	}, func(arg JobArg, res chan<- Result) error {
		a := arg.(Arg)
		r := a.X * a.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job3 := NewJob(Arg{
		X: 100, Y: 4,
	}, func(arg JobArg, res chan<- Result) error {
		a := arg.(Arg)
		r := a.X * a.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job4 := NewJob(Arg{
		X: 5, Y: 5,
	}, func(arg JobArg, res chan<- Result) error {
		a := arg.(Arg)
		r := a.X * a.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job5 := NewJob(Arg{
		X: 100, Y: 100,
	}, func(arg JobArg, res chan<- Result) error {
		a := arg.(Arg)
		r := a.X * a.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job6 := NewJob(Arg{
		X: 2, Y: 2,
	}, func(arg JobArg, res chan<- Result) error {
		a := arg.(Arg)
		r := a.X * a.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job7 := NewJob(Arg{
		X: 25, Y: 2,
	}, func(arg JobArg, res chan<- Result) error {
		a := arg.(Arg)
		r := a.X * a.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	job8 := NewJob(Arg{
		X: 10, Y: 2,
	}, func(arg JobArg, res chan<- Result) error {
		a := arg.(Arg)
		r := a.X * a.Y
		runIn := time.Duration(1000)
		fmt.Printf("job run in %d ms\n", runIn)
		time.Sleep(time.Millisecond * runIn)

		res <- r
		return nil
	})

	jobs := []Job{job1, job2, job3, job4, job5, job6, job7, job8}

	return jobs
}
func TestPoolRun(t *testing.T) {

	jobs := generateJobs()

	ctx, cancel := context.WithCancel(context.Background())
	defer func() { cancel() }()

	threadPool := NewThreadPool(4)
	threadPool.GenerateJobFrom(jobs)

	threadPool.Run(ctx)

	for {
		select {
		case r, ok := <-threadPool.Result():
			if !ok {
				break
			}

			i := r.(uint)
			if i <= 0 {
				t.Error("result is not valid")
			}
		case done := <-threadPool.Done:
			if !done {
				t.Error("worker is not done")
			}
			return
		}
	}
}

func TestPoolRunWithTimeout(t *testing.T) {

	jobs := generateJobs()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*800)
	defer func() { cancel() }()

	threadPool := NewThreadPool(4)
	threadPool.GenerateJobFrom(jobs)

	threadPool.Run(ctx)

	for {
		select {
		case r, ok := <-threadPool.Result():
			if !ok {
				break
			}

			i := r.(uint)
			if i <= 0 {
				t.Error("result is not valid")
			}
		case done := <-threadPool.Done:
			if !done {
				t.Error("worker is not done")
			}
			return
		}
	}
}
