package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
)

var zero = big.NewInt(0)
var one = big.NewInt(1)

func Factorial(n *big.Int) *big.Int {
	fact := big.NewInt(1)
	for n.Cmp(zero) > 0 {
		fact.Mul(fact, n)
		n.Sub(n, one)
	}

	return fact
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
