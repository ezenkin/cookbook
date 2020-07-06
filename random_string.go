package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println(generateRandomString(16))
	fmt.Println(generateRandomString2(16))
}

func generateRandomString(n int) string {
	buf := make([]byte, n)
	for i := range buf {
		buf[i] = letters[rand.Intn(len(letters))]
	}
	return string(buf[:n])
}

func generateRandomString2(n int) (string, error) {
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return "", err

	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}
