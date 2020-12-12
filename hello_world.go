// go lang hello world
package main

import "fmt"

func main() {
	fmt.Println("This shit is just like C. I'm gonna cry.")
}

// command line arguments
// echo prints its command line arguments.
import(
	"fmt"
	"os"
)
func main() {
	var a, b string
	for i:=1; i<len(os.Args); i++ {
		a += b + os.Args[i] //old string + space + next argument
		b = " "
	}
	fmt.Println(a)
}


// for loop (only loop in go lang)
// for initialisation; condition; post {
// 	statements
// }
// another form of echo.

import(
	"fmt"
	"os"
)

func main() {
	a, b := "", ""
	for _, arg := range os.Args[1:] {
		a += b + arg
		b = " "
	}
	fmt.Println(a)
}


// join function alternative
// to avoid the garbage of the old contents of a
import(
	"fmt"
	"os"
	"strings"
)
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// if we don't care much about output shape
func main() {
	fmt.Println(os.Args[1:])
}


// modifying to print from arg[0]
func main() {
	fmt.Println(os.Args[0:])
}

//modify to print index-value per line
import(
	"fmt"
	"os"
)
func main() {
	a, b := "", "\n"
	for i, arg:= range os.Args[0:] {
		a += b + string(i) + arg
		b = "\n"
	}
	fmt.Println(a)
}

// measure the difference in runtime between the 1st ways and using the Join function