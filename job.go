package tpool

import (
	"fmt"
)

// JobArg type
type JobArg interface{}

// JobFn type
type JobFn func(JobArg, chan<- Result) error

// Job type
type Job struct {
	arg JobArg
	op  JobFn
}

// NewJob the job constructor
func NewJob(arg JobArg, op JobFn) Job {
	return Job{
		arg: arg,
		op:  op,
	}
}

// Execute will execute job's operation
func (j Job) Execute(workerId uint, res chan<- Result) {
	fmt.Printf("execute job on worker = %d\n", workerId)
	j.op(j.arg, res)
}
