package rng

import (
	"math/rand"
	"sync"

	"github.com/seehuhn/mt19937"
)

type Mt19937 struct {
	mu   sync.Mutex
	rand *rand.Rand
	seed int64
}

func (m *Mt19937) Reseed(seed int64) {
	m.rand.Seed(seed)
	m.seed = seed
}

func NewRNG() *Mt19937 {
	m := rand.New(mt19937.New())

	return &Mt19937{
		rand: m,
	}
}

func (m *Mt19937) int() int {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.rand.Int()
}

func ShiftUint32(got, min, max uint32) uint32 {
	return got%(max-min+1) + min
}

func (m *Mt19937) Random(min uint32, max uint32) uint32 {
	return ShiftUint32(uint32(m.int()), min, max)
}

func (m *Mt19937) Seed() int64 {
	return m.seed
}
