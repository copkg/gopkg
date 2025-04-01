package util

import (
	"math/rand"
	"time"
)

func VerificationCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	captcha := make([]byte, length)

	for i := range captcha {
		captcha[i] = digits[rand.Intn(len(digits))]
	}

	return string(captcha)
}
