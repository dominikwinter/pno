package generator

import "fmt"

func GetFormat(formatFlag, date, country, gender, checksum string) (string, error) {
	switch formatFlag {
	case "1", "yymmddccgn":
		return GetShortDate(date) + country + gender + checksum, nil

	case "2", "yymmdd-ccgn":
		return GetShortDate(date) + "-" + country + gender + checksum, nil

	case "3", "yyyymmddccgn", "":
		return date + country + gender + checksum, nil

	case "4", "yyyymmdd-ccgn":
		return date + "-" + country + gender + checksum, nil

	default:
		return "", fmt.Errorf("unknown format '%s'", formatFlag)
	}
}
