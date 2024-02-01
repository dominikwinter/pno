package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/dominikwinter/pno/internal/val"
)

func Val(args []string) {
	fs := flag.NewFlagSet("pno val <pno>", flag.ExitOnError)

	helpFlag := fs.Bool("h", false, "Help")
	verboseFlag := fs.Bool("v", false, "Verbose")
	exitFlag := fs.Bool("e", false, "Return exit code 1 if not valid")

	fs.Parse(args)

	if *helpFlag || len(args) == 0 {
		fs.Usage()
		return
	}

	pno := args[len(args)-1]
	valid, _, err := val.ValidatePno(pno)

	if *verboseFlag {
		if err != nil {
			panic(err)
		}
	}

	if *exitFlag {
		if valid {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
		return
	} else {
		if valid {
			fmt.Println("valid")
		} else {
			fmt.Println("invalid")
		}
		return
	}
}
