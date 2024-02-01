package val

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/dominikwinter/pno/internal/gen"
)

var re = regexp.MustCompile(`(\d{2,4})(\d{2})(\d{2})-?(\d{2})(\d)(\d)`)

func ValidatePno(pno string) (bool, string, error) {
	results := re.FindAllStringSubmatch(pno, -1)

	if len(results) != 1 {
		return false, "", errors.New("invalid format")
	}

	matches := results[0]

	if len(matches) != 7 {
		return false, "", errors.New("invalid format")
	}

	year, err := fillYear(matches[1])

	if err != nil {
		return false, "", err
	}

	date := year + matches[2] + matches[3]
	country := matches[4]
	gender := matches[5]
	checksum := matches[6]

	genChecksum, err := gen.GetChecksum(date, country, gender)

	if err != nil {
		return false, "", err
	}

	return genChecksum == checksum, genChecksum, nil
}

func fillYear(year string) (string, error) {
	if len(year) == 2 {
		currentYear := time.Now().Year()
		year, err := strconv.Atoi(year)

		if err != nil {
			return "", fmt.Errorf("invalid year input: %w", err)
		}

		century := currentYear / 100

		if year < (currentYear % 100) {
			century++
		}

		return strconv.Itoa(century*100 + year), nil
	}

	return year, nil
}
