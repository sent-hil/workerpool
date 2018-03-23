# workerpool

workerpool is a simple library to have a pool of x workers to simultaneously call the given handler.

```go
// setup worker pool with handler
var numOfWorkers = 10
c, wg := New(numOfWorkers, func(x interface{}) error {
  // your handler goes here

  return nil
})

// queue jobs to be run by worker pool
for i := 0; i < 20; i++ {
  wg.Add(1)
  c <- i
}

// wait for all jobs to complete
wg.Wait()
```
