package main

import (
	"fmt"
	"io"
	"os"

	"github.com/davecgh/go-spew/spew"
	grammar "github.com/zenithar/nmap-service-peg"
)

func main() {
	// Read the input
	payload, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the input
	res, err := grammar.Parse("", payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	spew.Dump(res)
}
