package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateSubscriptionId generates a random string of length 6 to be used a subscription id
func GenerateSubscriptionId() string {
	// Define the length of the random string
	length := 6

	// Generate a byte slice of the desired length
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	// Encode the byte slice as a base64 string
	return base64.StdEncoding.EncodeToString(b)[:length]
}
