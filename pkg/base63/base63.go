package base63

import (
	"errors"
	"strings"
)

var ErrInvalidCharacters = errors.New("invalid character in string")

const (
	characterSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
	base         = 63
)

func Encode(num, minLen int) string {
	b := make([]byte, 0)
	p := 0
	for num > 0 {
		r := num % base
		num /= base
		b = append([]byte{characterSet[r]}, b...)
		p++
	}
	if p < minLen {
		for i := minLen - p; i > 0; i-- {
			b = append([]byte{characterSet[0]}, b...)
		}
	}

	return string(b)
}

func Decode(s string) (int, error) {
	var r, p int

	for i, v := range s {
		p = len(s) - (i + 1)
		pos := strings.IndexRune(characterSet, v)

		if pos == -1 {
			return 0, ErrInvalidCharacters
		}
		r += pos * pow(base, p)
	}

	return r, nil
}

func pow(base, p int) int {
	if p == 0 {
		return 1
	}
	return base * pow(base, p-1)
}
