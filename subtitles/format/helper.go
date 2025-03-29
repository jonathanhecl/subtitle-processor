// Package format provides functions for reading and writing different subtitle formats.
package format

import (
	"strconv"
	"strings"
	"unicode"
)

// toInt converts a string to an integer, ignoring errors.
// Used for parsing numeric values in subtitle formats.
func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// cleanText removes BOM and other special characters from text.
// This ensures consistent text processing across different subtitle formats.
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
