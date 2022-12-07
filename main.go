package main

import (
	"fmt"
	"os"

	"github.com/flp-wiki/faq/gif"
)

func main() {
	if err := gif.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
