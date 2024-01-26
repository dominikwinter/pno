package generator

import (
	"fmt"
	"strconv"
)

func GetShortDate(date string) string {
	return date[2:]
}

func GetChecksum(date, country, gender string) (string, error) {
	if l := len(date); l != 8 {
		return "", fmt.Errorf("date is not 8 characters long, but '%d'", l)
	}

	shortDate := GetShortDate(date)
	num := shortDate + country + gender
	numLen := len(num)

	var checksum int

	for i := numLen - 2; i >= 0; i -= 2 {
		n := num[i] - '0'
		checksum += int(n)
	}

	for i := numLen - 1; i >= 0; i -= 2 {
		n := (num[i] - '0') * 2

		if n > 9 {
			n -= 9
		}

		checksum += int(n)
	}

	return strconv.Itoa(checksum * 9 % 10), nil
}
