package rng

type RNG interface {
	Random(min uint32, max uint32) uint32
}
