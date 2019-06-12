package utils

import (
	"math/rand"
	"time"
)

func NewID() uint64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Uint64()
}
