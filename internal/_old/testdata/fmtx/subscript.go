package fmtx

import (
	"fmt"
)

type integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// Subscript renders n as a string using unicode subscript characters.
func Subscript[T integer](n T) string {
	digits := []rune(fmt.Sprint(n))

	for i, r := range digits {
		switch r {
		case '-':
			digits[i] = '₋'
		default:
			digits[i] = r - '0' + '₀'
		}
	}

	return string(digits)
}
