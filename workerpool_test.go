package woorkerpool

import (
	"sync"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("New", t, func() {
		s := sync.Map{}
		c, wg := New(10, func(x interface{}) error {
			time.Sleep(1 * time.Second)

			s.Store(x, true)
			return nil
		})

		startTime := time.Now()
		for i := 0; i < 20; i++ {
			wg.Add(1)
			c <- i
		}
		wg.Wait()
		endTime := time.Now()

		Convey("It should take ~2 seconds (20 jobs with 10 concurrency) to run", func() {
			So(startTime, ShouldNotHappenWithin, 1*time.Second, endTime)
			So(startTime, ShouldHappenWithin, 3*time.Second, endTime)
		})

		Convey("It should run given handler", func() {
			for i := 0; i < 20; i++ {
				_, ok := s.Load(i)
				So(ok, ShouldBeTrue)
			}
		})
	})
}
