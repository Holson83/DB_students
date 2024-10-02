package utilities

import (
	"strconv"
)

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}
