package generator

import (
	"fmt"
	"math/rand"
	"time"
)

const DATE_FORMAT = "20060102"

func GetRandomDate() (string, error) {
	now := time.Now()

	min := now.AddDate(-100, 0, 0).Unix()
	max := now.AddDate(-18, 0, 0).Unix()

	delta := max - min
	sec := rand.Int63n(delta) + min

	return time.Unix(sec, 0).Format(DATE_FORMAT), nil
}

func GetDate(date string) (string, error) {
	d, err := time.Parse(DATE_FORMAT, date)
	if err != nil {
		return "", fmt.Errorf("could not parse date: '%s': %w", date, err)
	}

	sec := d.Unix()

	return time.Unix(sec, 0).Format(DATE_FORMAT), nil
}
