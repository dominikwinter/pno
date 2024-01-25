package main

import (
	"testing"
)

func Test_getRandomDate(t *testing.T) {
	date := getRandomDate()

	if l := len(date); l != 8 {
		t.Errorf("Date is not 8 characters long, but %d", l)
	}
}

func Test_getRandomThreeDigits(t *testing.T) {
	three := getRandomThreeDigits(true)

	if l := len(three); l != 3 {
		t.Errorf("Three is not 3 characters long, but %d", l)
	}
}

func Test_getChecksum(t *testing.T) {
	date := "19800101"
	three := "123"

	checksum := getChecksum(&date, &three)

	if l := len(checksum); l != 1 {
		t.Errorf("Checksum is not 2 characters long, but %d", l)
	}
}
