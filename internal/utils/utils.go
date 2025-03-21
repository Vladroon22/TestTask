package utils

import (
	"regexp"
)

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$)")
	return emailRegex.MatchString(email)
}

func ValidatePhone(phone string) bool {
	phoneRegex := regexp.MustCompile(`^(?:\+7|8)?[\s-]?\(?\d{3}\)?[\s-]?\d{2,3}[\s-]?\d{2,3}[\s-]?\d{2,4}$`)
	return phoneRegex.MatchString(phone)
}
