package pgconcurrency

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

// main is the entry point for all Go programs.
func BuildRaceConditionFixed() {
	fmt.Println("Fixing RACE Conditions: Counter Value:", counter)
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	go incCounter2(1)
	go incCounter2(2)

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable.
func incCounter2(id int) {
	fmt.Println("Init Counter for Routine: ", id)

	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Capture the value of Counter.

		// Yield the thread and be placed back in queue.
		// “give the other goroutine a chance to run”

		atomic.AddInt64(&counter, 1)

		fmt.Println("Yield Thread Routine:", id)
		runtime.Gosched()
		fmt.Println("Running again Routine:", id)

	}
}
