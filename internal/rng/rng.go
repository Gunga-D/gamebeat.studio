//go:generate mockgen -destination=mocks/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package rng

type RNG interface {
	Random(min uint32, max uint32) uint32
}
