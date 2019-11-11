package counters_test

import (
	"counters"
	"fmt"
	"time"
)

//Synthetic use case for MaxTime, common i'm using it for measure function execution time.
func ExampleMaxTime() {
	//Run some steps of max time measurement
	for j := 0; j < 3; j++ {
		m := counters.MaxTime(0)
		for i := 0; i < 10; i++ {
			//Run some treads< and sleep time measurement
			go func() {
				start := time.Now().UnixNano()
				time.Sleep(time.Duration(i) * time.Duration(time.Millisecond))
				m.Store(time.Now().UnixNano() - start)
			}()
		}
		fmt.Println("Step ", j, " Max sleeping time: ", m.Get(), "ns")
		m.Reset()
	}
}
