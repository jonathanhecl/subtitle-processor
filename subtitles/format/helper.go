package format

import (
	"strconv"
)

func toInt(input string) (output int) {
	output, _ = strconv.Atoi(input)
	return output
}
