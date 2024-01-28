package gen

import (
	"fmt"
	"math/rand"
	"strconv"
)

func GetGenderNumber(gender string) (string, error) {
	var num int

	switch gender {
	case "m":
		// generate random single digit odd number
		num = rand.Intn(10) | 1 // force last bit to 1

	case "f":
		// generate random single digit even number
		num = rand.Intn(8)
		num += num & 1 // force last bit to 0

	default:
		return "", fmt.Errorf("unknown gender '%s'", gender)
	}

	return strconv.Itoa(num), nil
}

func GetRandomGenderNumber() (string, error) {
	if n := rand.Intn(2); n == 0 {
		return GetGenderNumber("m")
	} else {
		return GetGenderNumber("f")
	}
}
