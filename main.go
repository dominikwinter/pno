package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	date := getRandomDate()
	three := getRandomThreeDigits(true)
	checksum := getChecksum(&date, &three)

	fmt.Println(date + three + checksum)
}

func getRandomDate() string {
	now := time.Now()

	min := now.AddDate(-100, 0, 0).Unix()
	max := now.AddDate(-18, 0, 0).Unix()

	delta := max - min
	sec := rand.Int63n(delta) + min

	return time.Unix(sec, 0).Format("20060102")
}

func getRandomThreeDigits(male bool) string {
	two := rand.Intn(99)
	one := rand.Intn(9)

	var last int

	if male {
		if one%2 == 1 {
			last = one
		} else {
			last = (one + 1) % 10
		}
	} else {
		if one%2 == 0 {
			last = one
		} else {
			last = (one + 1) % 10
		}
	}

	return fmt.Sprintf("%02d%d", two, last)
}

func getChecksum(date *string, three *string) string {
	short_date := (*date)[len(*date)-6:]
	short_checksum := short_date + *three

	sum := 0
	m := 2

	for i := 0; i < len(short_checksum); i++ {
		num, _ := strconv.Atoi(string(short_checksum[i]))
		num = num * m

		if num >= 10 {
			sum += (num % 10) + 1
		} else {
			sum += num
		}

		if m == 2 {
			m = 1
		} else {
			m = 2
		}
	}

	checksum := 10 - (sum % 10)

	if checksum == 10 {
		checksum = 0
	}

	return strconv.Itoa(checksum)
}
