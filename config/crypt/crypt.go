package crypt

import "strings"

func Encode(pw string) string {
	replacedText := strings.ReplaceAll(pw, "a", "*")
	replacedText = strings.ReplaceAll(replacedText, "i", "%")
	replacedText = strings.ReplaceAll(replacedText, "u", "$")
	replacedText = strings.ReplaceAll(replacedText, "e", "@")
	replacedText = strings.ReplaceAll(replacedText, "o", "!")

	return replacedText
}

func Decode(pw string) string {
	replacedText := strings.ReplaceAll(pw, "*", "a")
	replacedText = strings.ReplaceAll(replacedText, "%", "i")
	replacedText = strings.ReplaceAll(replacedText, "$", "u")
	replacedText = strings.ReplaceAll(replacedText, "@", "e")
	replacedText = strings.ReplaceAll(replacedText, "!", "o")

	return replacedText
}