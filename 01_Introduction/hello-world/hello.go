package main // package to group functions

import (
	"fmt" // import fmt standard package for formatting text, including print to console
	"rsc.io/quote"
)

// main function is executed by default when running the main package
func main() {
	fmt.Println(quote.Go())
}

/* Result:
$ go run .
Hello World
*/
