package stringformatter

import "strings"

func ReplaceI18nLetters(s string) string {
	str := s

	str = strings.ReplaceAll(str, "é", "e")
	str = strings.ReplaceAll(str, "ö", "o")
	str = strings.ReplaceAll(str, "å", "a")
	str = strings.ReplaceAll(str, "æ", "ae")
	str = strings.ReplaceAll(str, "í", "i")

	return str
}
