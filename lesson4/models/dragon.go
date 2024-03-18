package models

import (
	"math/rand"
)

type Dragon struct {
	Heads int
}

func NewDragon() *Dragon {
	return &Dragon{Heads: startHeads(50, 150)}
}

func (h *Dragon) IsLoose() bool {
	return h.Heads <= 0
}

func (h *Dragon) IsWin() bool {
	return h.Heads >= 200
}

func startHeads(min, max int) int {
	return min + rand.Intn(max-min)
}
