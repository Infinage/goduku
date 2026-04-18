package core

import (
	"iter"
	"math/rand/v2"
)

// Range is inclusive of both ends
func sequence(start, end uint8, shuffle bool) iter.Seq[uint8] {
	return func(yield func(uint8) bool) {
		if !shuffle {
			for i := start; i <= end; i++ {
				if !yield(i) {
					return
				}
			}
			return
		}

		// Unfortunately need to allocate for shuffling
		size := int(end - start + 1)
		rng := make([]uint8, size)
		for i := range size {
			rng[i] = start + uint8(i)
		}

		rand.Shuffle(len(rng), func(i, j int) {
			rng[i], rng[j] = rng[j], rng[i]
		})

		for _, n := range rng {
			if !yield(n) {
				return
			}
		}
	}
}
