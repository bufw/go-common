package bytesize

import (
	"fmt"
	"strconv"
)

//byte size
const (
	B  uint64 = 1
	KB uint64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

//ByteSizeToString to string
func ByteSizeToString(length uint64, decimals int) (out string) {
	var unit string
	var i uint64
	var remainder uint64

	if length > TB {
		unit = "TB"
		i = length / TB
		remainder = length - (i * TB)
	} else if length > GB {
		unit = "GB"
		i = length / GB
		remainder = length - (i * GB)
	} else if length > MB {
		unit = "MB"
		i = length / MB
		remainder = length - (i * MB)
	} else if length > KB {
		unit = "KB"
		i = length / KB
		remainder = length - (i * KB)
	} else {
		return strconv.FormatUint(length, 10) + " B"
	}

	if decimals == 0 {
		return strconv.FormatUint(i, 10) + " " + unit
	}

	width := 0
	if remainder > GB {
		width = 12
	} else if remainder > MB {
		width = 9
	} else if remainder > KB {
		width = 6
	} else {
		width = 3
	}

	remainderString := strconv.FormatUint(remainder, 10)
	for iter := len(remainderString); iter < width; iter++ {
		remainderString = "0" + remainderString
	}
	if decimals > len(remainderString) {
		decimals = len(remainderString)
	}

	return fmt.Sprintf("%d.%s %s", i, remainderString[:decimals], unit)
}
