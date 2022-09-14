package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

const workers = 5

func Factorial(n uint64) uint64 {
	chunk := n/workers + 1
	out := make(chan uint64)
	wg := sync.WaitGroup{}

	for i := uint64(1); i <= n; i += chunk {
		lower := i
		upper := i + chunk - 1
		if upper > n {
			upper = n
		}

		wg.Add(1)
		go func(l, u uint64) {
			part := rec(u, l)
			out <- part
			wg.Done()
		}(lower, upper)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	f := uint64(1)
	for part := range out {
		f *= part
	}
	return f
}

func rec(n uint64, m uint64) uint64 {
	if n <= m {
		return m
	}

	return n * rec(n-1, m)
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
