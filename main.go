package main

import (
	"fmt"
	"os"

	"github.com/dominikwinter/pno/cmd"
)

func main() {
	const usage = `Usage of pno <command> [<args>]:

  commands:
    gen     Generate a personal number

  args:
    -h      Help`

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)

		os.Exit(1)
	}

	command, args := os.Args[1], os.Args[2:]

	switch command {
	case "gen":
		cmd.Run(args)
	default:
		fmt.Fprintf(os.Stderr, "Unrecognized command %s.\n\n%s\n", command, usage)
		os.Exit(1)
	}
}
