package main

import (
	"os"

	"github.com/digyx/Turnkey/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
