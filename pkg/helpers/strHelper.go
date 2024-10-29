package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// StringSliceToString converts a slice of strings to a comma-delimited string
func StringSliceToString(src []string) (tgt string) {
	tgt = "["
	for i := 0; i < len(src); i++ {
		joinChr := ", "
		if i == 0 {
			joinChr = ""
		}
		tgt = fmt.Sprintf("%s%s\"%s\"", tgt, joinChr, src[i])
	}
	tgt = fmt.Sprintf("%s]", tgt)
	return
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if n < 0 {
		return "", fmt.Errorf("n must be positive: %d", n)
	}
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

// Contains checks if a string is present in a slice.
// It will return true if the string is present, false
// if it is not.
func Contains(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// ToTitleCase converts the first character of each word to upper case and the rest to lower case.
func ToTitleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	words := strings.Fields(s)
	for i, word := range words {
		r := strings.ToLower(word)
		words[i] = strings.ToUpper(string(r[0])) + r[1:]
	}
	return strings.Join(words, " ")
}
