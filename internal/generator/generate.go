package generator

import "fmt"

func GeneratePno(dateFlag, countryFlag, genderFlag, formatFlag string) (string, error) {
	date, err := GenerateStr(dateFlag, GetDate, GetRandomDate)
	if err != nil {
		return "", fmt.Errorf("failed to get date: %w", err)
	}

	country, err := GenerateStr(countryFlag, GetCountry, GetRandomCountry)
	if err != nil {
		return "", fmt.Errorf("failed to get country: %w", err)
	}

	gender, err := GenerateStr(genderFlag, GetGenderNumber, GetRandomGenderNumber)
	if err != nil {
		return "", fmt.Errorf("failed to get gender: %w", err)
	}

	checksum, err := GetChecksum(date, country, gender)
	if err != nil {
		return "", fmt.Errorf("failed to get checksum: %w", err)
	}

	return GetFormat(formatFlag, date, country, gender, checksum)
}

func GenerateStr[I string, O string](
	flag I,
	getFunc func(I) (O, error),
	getRandomFunc func() (O, error),
) (O, error) {
	if len(flag) > 0 {
		return getFunc(flag)
	}

	return getRandomFunc()
}
