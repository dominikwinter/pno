package gen

import (
	"strconv"
	"testing"
)

func Test_GetRandomDate(t *testing.T) {
	for i := 0; i < 1000; i++ {
		date, err := GetRandomDate()
		if err != nil {
			t.Error(err)
		}

		if l := len(date); l != 8 {
			t.Errorf("Date is not 8 characters long, but %d", l)
		}
	}
}

func Test_getRandomBirthCountry(t *testing.T) {
	for i := 0; i < 1000; i++ {
		country, err := GetRandomCountry()
		if err != nil {
			t.Error(err)
		}

		if l := len(country); l != 2 {
			t.Errorf("Country is not 2 characters long, but %d", l)
		}
	}
}

func Test_getRandomGenderNumber(t *testing.T) {
	for i := 0; i < 1000; i++ {
		gender, err := GetRandomGenderNumber()
		if err != nil {
			t.Error(err)
		}

		if l := len(gender); l != 1 {
			t.Errorf("Gender is not 1 characters long, but %d", l)
		}
	}
}

func Test_GetChecksum(t *testing.T) {
	date := "19800101"
	country := "12"
	gender := "1"

	checksum, err := GetChecksum(date, country, gender)
	if err != nil {
		t.Error(err)
	}

	if l := len(checksum); l != 1 {
		t.Errorf("Checksum is not 1 characters long, but %d", l)
	}
}

func Test_GetGenderNumber(t *testing.T) {
	t.Run("male", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			str, err := GetGenderNumber("m")
			if err != nil {
				t.Error(err)
			}

			num, err := strconv.Atoi(str)
			if err != nil {
				t.Error(err)
			}

			if num%2 == 0 {
				t.Errorf("%d is not a male random number", num)
			}
		}
	})

	t.Run("female", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			str, err := GetGenderNumber("f")
			if err != nil {
				t.Error(err)
			}

			num, err := strconv.Atoi(str)
			if err != nil {
				t.Error(err)
			}

			if num%2 != 0 {
				t.Errorf("%d is not a female random number", num)
			}
		}
	})
}
