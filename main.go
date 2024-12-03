package main

import (
	"fmt"

	"playground/datastructures"

	pgconcurrency "playground/pg_concurrency"

	"rsc.io/quote/v4"
)

func init() {
	fmt.Println("Init call...")
}

func main() {
	printNumbers()
	fmt.Println("Using Quote: " + quote.Go())
	testingDataStructures()
	pgconcurrency.Run()
	pgconcurrency.RescheduleRunner()
	pgconcurrency.BuildRaceCondition()

}

func printNumbers() {
	fmt.Printf("Printing  %f  \n", 12.0)
}

func testingDataStructures() {
	datastructures.TestSlices()
}
