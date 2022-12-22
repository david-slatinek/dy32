package random

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func random(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func String(min, max int) string {
	return random(rand.Intn(max-min+1) + min)
}

func Int(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func Float(min, max int) float64 {
	return rand.Float64() * float64((max-min+1)+min)
}
