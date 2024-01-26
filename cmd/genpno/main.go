package main

import (
	"flag"
	"fmt"

	"github.com/dominikwinter/genpno/internal/generator"
)

func main() {
	helpFlag := flag.Bool("h", false, "Help")
	verboseFlag := flag.Bool("v", false, "Verbose")

	dateFlag := flag.String("d", "", "Date (format: yyyymmdd) (default random)")
	countryFlag := flag.String("c", "", "Country code (default random)")
	genderFlag := flag.String("g", "", "Gender (m/f) (default random)")
	formatFlag := flag.String("f", "", `Output forma (default 3):
	1. yymmddccgn
	2. yymmdd-ccgn
	3. yyyymmddccgn
	4. yyyymmdd-ccgn`)

	flag.Parse()

	if *helpFlag {
		flag.Usage()
		return
	}

	pno, err := generator.GeneratePno(*dateFlag, *countryFlag, *genderFlag, *formatFlag)

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
