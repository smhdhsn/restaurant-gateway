package helper

import (
	"strconv"
)

// StrToUint is the funcion responsible for converting a string typed number to a 32 bit uint.
func StrToUint(str string) (uint32, error) {
	num, err := strconv.ParseUint(str, 10, 32)
	return uint32(num), err
}
