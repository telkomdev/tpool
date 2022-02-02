package tpool

// Waiter the waiter interface
type Waiter interface {
	Done()
	Wait()
	Close()
}

type empty struct{}

type waiter struct {
	Capacity uint
	Sig      chan empty
}

func newWaiter(c uint) Waiter {
	sig := make(chan empty, c)
	return &waiter{Capacity: c, Sig: sig}
}

func (w *waiter) Done() {
	e := empty{}
	w.Sig <- e
}

func (w *waiter) Wait() {
	var i uint
	for i = 0; i < w.Capacity; i++ {
		<-w.Sig
	}
}

func (w *waiter) Close() { close(w.Sig) }
