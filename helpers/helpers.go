package helpers

import (
	"math/rand"
)

type Meds struct {
	Age         int
	Experience  int
	Speciaility string
}

func RandomNum(n int) int {
	num := rand.Intn(n)
	return num
}
