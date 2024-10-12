package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	//Adding 2 new logical processors for each goroutine (counting semaphores for the goroutines)
	wg.Add(2)

	go decrementCounter(1)
	go decrementCounter(3)

	//blocks the main thread so that the main func can not return unless the goroutines are done
	wg.Wait()
	fmt.Println("Final Count: ", counter)
}
func decrementCounter(id int) {
	// makes sure that the function decrements the semaphore when done and tell the main the goroutine is done
	defer wg.Done()
	// According to this loop the value of count should not be as the result of running this program due to race conditions
	//in accessing the counter which is a shared memory with different parallel executing threads(logical processors / os threads)
	// Since the copy of the counter is used and not a reference to the object being modified by parallel execution then the
	//next process that receives the value of counter will be different from the one that is running and updating the value of counter.
	for count := id; count > 0; count-- {
		value := counter

		//Yielding the thread and making it to be placed back in queue.
		runtime.Gosched()

		value--

		counter = value

	}
}
