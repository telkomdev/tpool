package tpool

import (
	"context"
	"fmt"
)

type Result interface{}

func worker(id uint, ctx context.Context,
	waiter Waiter, jobs <-chan Job, res chan<- Result) {
	defer func() { waiter.Done() }()

	for {

		// The try-receive operation here is to
		// try to exit the worker goroutine as
		// early as possible. Try-receive
		// optimized by the standard Go
		// compiler, so they are very efficient.
		select {
		case <-ctx.Done():
			fmt.Printf("worker canceled\n")
			return
		default:
		}

		// Even if ctx.Done() is set to closed, the first
		// branch in this select block might be
		// still not selected for some loops
		// so the try-receive operation above is essential
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}

			job.Execute(id, res)

		case <-ctx.Done():
			fmt.Printf("worker canceled\n")
			return
		}
	}
}

// ThreadPool type
type ThreadPool struct {
	n      uint
	result chan Result
	jobs   chan Job
	Done   chan bool
}

// NewThreadPool the ThreadPool's constructor
func NewThreadPool(n uint) ThreadPool {
	result := make(chan Result, n)
	jobs := make(chan Job, n)
	done := make(chan bool, 1)
	return ThreadPool{
		n:      n,
		result: result,
		jobs:   jobs,
		Done:   done,
	}
}

// Result will return result chan
func (t ThreadPool) Result() <-chan Result { return t.result }

// Jobs will return jobs chan
func (t ThreadPool) Jobs() chan<- Job { return t.jobs }

// GenerateJobFrom a function helper for generate job from job lists
func (t ThreadPool) GenerateJobFrom(jobs []Job) {
	go func() {
		for _, j := range jobs {
			t.jobs <- j
		}

		close(t.jobs)
	}()
}

// Run will run thread pool
func (t ThreadPool) Run(ctx context.Context) {

	go func() {
		waiter := newWaiter(t.n)
		defer func() { waiter.Close() }()

		var i uint
		for i = 0; i < t.n; i++ {

			go worker(i, ctx, waiter, t.jobs, t.result)
		}

		waiter.Wait()
		t.Done <- true
		close(t.result)
	}()
}
