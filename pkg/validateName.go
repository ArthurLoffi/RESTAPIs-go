package pkg

import "regexp"

var nameRegex = regexp.MustCompile(`^[a-zA-ZÀ-ÿ\s]+$`)

func ValidateName(s string) bool {
	return nameRegex.MatchString(s)
}