package woorkerpool

import "sync"

func New(numWorkers int, handler func(interface{}) error) (chan interface{}, *sync.WaitGroup) {
	c := make(chan interface{})
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		go func() {
			for cc := range c {
				if err := handler(cc); err != nil {
					panic(err)
				}
				wg.Done()
			}
		}()
	}

	return c, &wg
}
