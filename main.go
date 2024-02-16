package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/dominikwinter/pno/cmd"
)

func main() {
	const usage = `Usage of pno <command> [<args>]

  commands:
    gen     Generate a personal number
    val     Validate a personal number

  args:
    -h      Help
    -V      Version`

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)

		os.Exit(1)
	}

	command, args := os.Args[1], os.Args[2:]

	switch command {
	case "-h":
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(0)

	case "-V":
		info, ok := debug.ReadBuildInfo()
		if !ok {
			return
		}

		var revision string
		var time string

		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				revision = setting.Value
			}
			if setting.Key == "vcs.time" {
				time = setting.Value
			}
		}

		fmt.Fprintf(os.Stdout, "%s %s\n", revision, time)
		os.Exit(0)

	case "gen":
		cmd.Gen(args)

	case "val":
		cmd.Val(args)

	default:
		fmt.Fprintf(os.Stderr, "Unrecognized command %s.\n\n%s\n", command, usage)
		os.Exit(1)
	}
}
