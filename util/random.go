package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a rondom integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sub strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sub.WriteByte(c)
	}

	return sub.String()
}

func RandomOwner() string {
	return RandomString(6)
}
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com",RandomString(6))
}

func RandomMoney() int64 {
	return RandomInt(10, 10000)
}

func RandomCurrency() string {
	currencies := []string{USD, EUR, RFW,CAD}
	return currencies[rand.Intn(len(currencies))]
}
