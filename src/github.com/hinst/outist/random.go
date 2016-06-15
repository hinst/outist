package outist

import (
	"math/rand"
	"time"
)

func CreateRandom() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
