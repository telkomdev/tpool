package tpool

import (
	"fmt"
)

// JobFn type
type JobFn[T any, R any] func(T, chan<- R) error

// Job type
type Job[T any, R any] struct {
	arg T
	op  JobFn[T, R]
}

// NewJob the job constructor
func NewJob[T any, R any](arg T, op JobFn[T, R]) Job[T, R] {
	return Job[T, R]{arg: arg, op: op}
}

// Execute will execute job's operation
func (j Job[T, R]) Execute(workerId uint, res chan<- R) {
	fmt.Printf("execute job on worker = %d\n", workerId)
	j.op(j.arg, res)
}
