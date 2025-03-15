package format

import (
	"strconv"
	"strings"
	"unicode"
)

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func cleanText(text string) string {
	// Remove BOM and other special characters
	text = strings.TrimPrefix(text, "\ufeff")
	text = strings.TrimPrefix(text, "\xef\xbb\xbf")
	
	// Remove other potential Unicode marks
	text = strings.Map(func(r rune) rune {
		if unicode.IsControl(r) && r != '\n' && r != '\r' {
			return -1
		}
		return r
	}, text)
	
	return strings.TrimSpace(text)
}
