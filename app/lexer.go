package app

import (
	"regexp"
)

// The individual lex functions return an empty string if they failed to read a
// token.

var digitRegex = regexp.MustCompile("[0-9]")

func lexNumber(input string, i int) string {
	token := ""

	for j := i; j < len(input); j++ {
		char := input[j : j+1]

		if digitRegex.MatchString(char) {
			token += char
		} else {
			return ""
		}
	}

	return token
}

func lexBoolean(input string, i int) string {
	if i+5 <= len(input) && input[i:i+5] == "false" {
		return "false"
	}

	if i+4 <= len(input) && input[i:i+4] == "true" {
		return "true"
	}

	return ""
}

func lexString(input string, i int) string {
	return ""
}

func Lex(input string) []string {
	tokens := []string{}
	i := 0

	check := func(token string) bool {
		if len(token) == 0 {
			return false
		}

		tokens = append(tokens, token)
		i += len(token)
		return true
	}

	for i < len(input) {
		var token string

		token = lexNumber(input, i)

		if check(token) {
			continue
		}

		token = lexBoolean(input, i)

		if check(token) {
			continue
		}

		break
	}

	return tokens
}