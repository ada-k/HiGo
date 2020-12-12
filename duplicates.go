package main

import(
	"fmt"
	"os"
	"bufio"
)

// prints each line that appears more than once in standardinput
func main() {
	counts:= make(map[string]int)
	input:= bufio.NewScanner(os.Stdin) // standard input
	for input.Scan() {
		counts[input.Text()]++

	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}

	}

}

