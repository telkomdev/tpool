package tpool

import (
	"context"
	"fmt"
)

func worker[T any, R any](id uint, ctx context.Context,
	waiter Waiter, jobs <-chan Job[T, R], res chan<- R) {
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
type ThreadPool[T any, R any] struct {
	n      uint
	result chan R
	jobs   chan Job[T, R]
	Done   chan bool
}

// NewThreadPool the ThreadPool's constructor
func NewThreadPool[T any, R any](n uint) ThreadPool[T, R] {
	result := make(chan R, n)
	jobs := make(chan Job[T, R], n)
	done := make(chan bool, 1)
	return ThreadPool[T, R]{
		n:      n,
		result: result,
		jobs:   jobs,
		Done:   done,
	}
}

// Result will return result chan
func (t ThreadPool[T, R]) Result() <-chan R { return t.result }

// Jobs will return jobs chan
func (t ThreadPool[T, R]) Jobs() chan<- Job[T, R] { return t.jobs }

// GenerateJobFrom a function helper for generate job from job lists
func (t ThreadPool[T, R]) GenerateJobFrom(jobs []Job[T, R]) {
	go func() {
		for _, j := range jobs {
			t.jobs <- j
		}

		close(t.jobs)
	}()
}

// Run will run thread pool
func (t ThreadPool[T, R]) Run(ctx context.Context) {

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
