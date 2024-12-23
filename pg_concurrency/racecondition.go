package pgconcurrency

import (
	"fmt"
	"runtime"
)

var (
	// counter is a variable incremented by all goroutines.
	counter int64
)

// main is the entry point for all Go programs.
func BuildRaceCondition() {
	fmt.Println("Building RACE Conditions")
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	go incCounter(1)
	go incCounter(2)

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable.
func incCounter(id int) {
	fmt.Println("Init Counter for Routine: ", id)

	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Capture the value of Counter.
		value := counter

		// Yield the thread and be placed back in queue.
		// “give the other goroutine a chance to run”

		fmt.Println("Yield Thread Routine:", id)
		runtime.Gosched()
		fmt.Println("Running again Routine:", id)

		// Increment our local value of Counter.
		value++

		// Store the value back into Counter.
		counter = value
		fmt.Printf("Counter Value %d for Routine: %d\n", counter, id)
	}
}
