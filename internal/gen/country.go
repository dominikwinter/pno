package gen

import (
	"fmt"
	"math/rand"
	"strconv"
)

func GetCountry(country string) (string, error) {
	if l := len(country); l != 2 {
		return "", fmt.Errorf("country is not 2 characters long, but '%d'", l)
	}

	c, err := strconv.Atoi(country)
	if err != nil {
		return "", fmt.Errorf("could not parse country '%s': %w", country, err)
	}

	return fmt.Sprintf("%02d", c), nil
}

func GetRandomCountry() (string, error) {
	country := rand.Intn(99)

	return fmt.Sprintf("%02d", country), nil
}
