package hello

import (
	// quoteV3 "rsc.io/quote/v3"
	"rsc.io/quote/v3"
)

// Hello test
func Hello() string {
	return quote.HelloV3()
}

// Proverb test
func Proverb() string {
	return quote.Concurrency()
}
