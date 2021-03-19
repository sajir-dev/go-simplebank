package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt to generate a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString to generate a random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner to generate random owner name for account
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount between 0 and 100
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency to generate random currency for the account
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	k := len(currencies)
	return currencies[rand.Intn(k)]
}
