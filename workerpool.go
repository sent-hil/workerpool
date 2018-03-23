package woorkerpool

import "sync"

// New spawns a pool of numWorkers workers to simultaneously call the given
// handler.
//
// It returns the a chan that's used to queue the jobs and sync.WaitGroup to
// indicate when all jobs are done processing.
func New(numWorkers int, handler func(interface{}) error) (chan interface{}, *sync.WaitGroup) {
	var (
		inputCh = make(chan interface{})
		waiter  sync.WaitGroup
	)

	for i := 0; i < numWorkers; i++ {
		go func() {
			for input := range inputCh {
				if err := handler(input); err != nil {
					panic(err)
				}
				waiter.Done()
			}
		}()
	}

	return inputCh, &waiter
}
