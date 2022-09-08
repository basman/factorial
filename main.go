package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("missing argument <N>")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("invalid number '%v'\n", os.Args[1])
	}

	f := Factorial(uint64(n))

	fmt.Printf("%v! = %v\n", n, f)
}
