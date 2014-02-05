// example_command - shows the basic usage of mango
//
// Description:
//
// It doesn't take much to create manual pages with mango. Just put a comment
// at the top of your source code file and write down stuff you want to include
// in the manual page like this. Feel free to add as many sections as you want.
package main

import (
	"flag"
)

var (
	optFoo = flag.Bool("foo", false, "this text should show up in the manual page")

	// If the flag definition follows a comment like this, mango will use
	// the text in the comment as description in the manual page.
	optBar = flag.Bool("bar", false, "the above comment should show up in the manual page")
	optBaz = false
)

func init() {
	// These two calls reference the same variable and should appear
	// grouped in the manual page.
	flag.Bool(&optBaz, "baz", false, "two calls, but one entry in the manual")
	flag.Bool(&optBaz, "b", false, "two calls, but one entry in the manual")
}

func main() {
	return
}
