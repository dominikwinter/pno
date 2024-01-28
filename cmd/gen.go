package cmd

import (
	"flag"
	"fmt"

	"github.com/dominikwinter/pno/internal/gen"
)

func Run(args []string) {
	fs := flag.NewFlagSet("pno gen", flag.ExitOnError)

	helpFlag := fs.Bool("h", false, "Help")
	verboseFlag := fs.Bool("v", false, "Verbose")

	dateFlag := fs.String("d", "", "Date (format: yyyymmdd) (default random)")
	countryFlag := fs.String("c", "", "Country code (default random)")
	genderFlag := fs.String("g", "", "Gender (m/f) (default random)")
	formatFlag := fs.String("f", "", `Output forma (default 3):
	1. yymmddccgn
	2. yymmdd-ccgn
	3. yyyymmddccgn
	4. yyyymmdd-ccgn`)

	fs.Parse(args)

	if *helpFlag {
		fs.Usage()
		return
	}

	pno, err := gen.GeneratePno(*dateFlag, *countryFlag, *genderFlag, *formatFlag)

	if err != nil {
		if *verboseFlag {
			panic(err)
		} else {
			fmt.Println(err)
			return
		}
	}

	fmt.Println(pno)
}
