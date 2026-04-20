package main

import (
	"math/rand/v2"
	"strings"
)

func GenerateRandomLinkCode(size uint) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var shortCode strings.Builder
	for i := 0; i <= int(size); i++ {
		shortCode.WriteString(string(alphabet[rand.IntN(len(alphabet))]))
	}

	return shortCode.String()
}

