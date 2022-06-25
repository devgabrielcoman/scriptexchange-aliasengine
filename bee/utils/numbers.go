package utils

import "strconv"

func LenientAtoi(stringDate string) int {
	num, err := strconv.Atoi(stringDate)
	if err == nil {
		return num
	} else {
		return 0
	}
}

func LenientAtoi64(date string) int64 {
	return int64(LenientAtoi(date))
}
