package helpers

import (
	"html"
	"unicode"
)

func TruncateByte(s string, maxBytes int) string {
	if len(s) <= maxBytes { // len() returns byte count
		return s
	}
	return s[:maxBytes]
}

func SanitizeString(s string) string {
	// Создаем слайс рун с запасом по емкости
	result := make([]rune, 0, len(s))

	for _, r := range s {
		// unicode.IsPrint проверяет, является ли символ печатным
		// (буквы, цифры, пунктуация, пробелы)
		// Также отсекает управляющие символы типа \x00, \x1F и т.д.
		if unicode.IsPrint(r) || r == '\n' || r == '\t' {
			result = append(result, r)
		}
	}
	return string(result)
}

func EscapeString(text string, maxLen int) string {
	return html.EscapeString(TruncateByte((SanitizeString(text)), maxLen))
}
