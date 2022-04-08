package tpool

import (
	"fmt"
)

// JobArg type
type JobArg interface{}

// JobFn type
type JobFn[T any] func(T, chan<- Result) error

// Job type
type Job[T any] struct {
	arg T
	op  JobFn[T]
}

// NewJob the job constructor
func NewJob[T any](arg T, op JobFn[T]) Job[T] {
	return Job[T]{arg: arg, op: op}
}

// Execute will execute job's operation
func (j Job[T]) Execute(workerId uint, res chan<- Result) {
	fmt.Printf("execute job on worker = %d\n", workerId)
	j.op(j.arg, res)
}
