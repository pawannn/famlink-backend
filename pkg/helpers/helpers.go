package helpers

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenerateRandom() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rand.Intn(900000) + 100000
	return randomNumber
}

func GenerateUUID() string {
	return uuid.New().String()
}
