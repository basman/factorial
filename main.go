package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func Factorial(n uint64) uint64 {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("missing argument <N>")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("invalid number '%v'\n", os.Args[1])
	}

	start := time.Now()

	// compute factorial
	f := Factorial(uint64(n))

	fmt.Printf("computation took %v ns\n", time.Now().Sub(start).Nanoseconds())

	fmt.Printf("\n%v! = %v\n", n, f)
}
