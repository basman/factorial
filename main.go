package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"sync"
	"time"
)

const numWorkers = 6

var zero = big.NewInt(0)
var one = big.NewInt(1)

func Factorial(n *big.Int) *big.Int {
	nWorker := big.NewInt(numWorkers)
	nCompRemainder := big.NewInt(0)
	nComp := big.NewInt(0)

	nComp.DivMod(n, nWorker, nCompRemainder)

	out := make(chan *big.Int)
	wg := sync.WaitGroup{}

	// launch workers
	iWorker := big.NewInt(0).Set(nWorker)
	for iWorker.Cmp(zero) > 0 && nComp.Cmp(zero) > 0 {
		start := big.NewInt(0)
		count := big.NewInt(0)

		iWorker.Sub(iWorker, one) // decrease

		count.Set(nComp)
		start.Mul(iWorker, nComp)

		wg.Add(1)
		go work(start, count, out, &wg)
	}

	// launch one worker for remainder
	if nCompRemainder.Cmp(zero) > 0 {
		start := big.NewInt(0)
		start.Mul(nWorker, nComp)

		wg.Add(1)
		go work(start, nCompRemainder, out, &wg)
	}

	// signal end of work to collector
	go func() {
		wg.Wait()
		close(out)
	}()

	// collect
	fact := big.NewInt(1)
	for m := range out {
		fact.Mul(fact, m)
	}

	return fact
}

func work(start *big.Int, count *big.Int, out chan *big.Int, wg *sync.WaitGroup) {
	origStart := big.NewInt(0)
	origStart.Set(start)

	fact := big.NewInt(1)

	for count.Cmp(zero) > 0 {
		start.Add(start, one) // next value
		fact.Mul(fact, start) // actual work
		count.Sub(count, one) // countdown
	}

	out <- fact
	wg.Done()
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
	f := Factorial(big.NewInt(int64(n)))

	fmt.Printf("%v! uses %v bits for storage (%v bytes)\n", n, f.BitLen(), len(f.Bytes()))
	fmt.Printf("computation took %v µs\n", time.Now().Sub(start).Microseconds())

	fmt.Printf("\n%v! = %v\n", n, f)
}
