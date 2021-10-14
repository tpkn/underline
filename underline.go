package underline

import (
	"strings"
	"unicode/utf8"
)

// Solid makes solid underline
func Solid(text string, with_text bool) string {
	return Custom(text, "─", with_text)
}

// Dashed makes dashed underline
func Dashed(text string, with_text bool) string {
	return Custom(text, "-", with_text)
}

// Dotted makes dotted underline
func Dotted(text string, with_text bool) string {
	return Custom(text, ".", with_text)
}

// Custom makes custom style underline
func Custom(text, symbol string, with_text bool) string {
	under := make([]string, utf8.RuneCountInString(text))
	for i := range under {
		under[i] = symbol
	}
	
	result := strings.Join(under, "")
	
	if with_text {
		result = text + "\n" + result
	}
	
	return result
}
