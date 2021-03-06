// Package shuffler creates slices of integers
// random, but reproducible; based on the ID of the user;
// classes of users see the same random order each time they visit;
// each page has a different randomization of appropriate length.
package shuffler

import (
	"log"
	"math/rand"
)

type shufflerT struct {
	Seed       int // Seed for shuffling; typically the user ID
	Variations int // Seed modulo Variations is the actual seed. Determines how many different shuffled sets are derived from various seeds, before patterns repeat

	MaxElements int // The number of elements to shuffle; typically the largest number of input groups across all pages of a questionnaire.
}

// New creates a Shuffler for creating deterministic variations
// of a slice []int{1,2,3...,MaxElements}
//
// seed is the seed for the randomizer.
// Variations is the number of classes.
// MaxElements is the slice length.
func New(seed int, variations int, maxNumberOfElements int) *shufflerT {
	s := shufflerT{}
	s.Variations = variations

	s.MaxElements = maxNumberOfElements
	s.Seed = seed
	return &s
}

// Slice generates a shuffled slice.
// Param iter gives the number of shufflings;
// iter is reduced to its modulo 7 - performance
func (s *shufflerT) Slice(shufflingRepetitions int) []int {

	order := make([]int, s.MaxElements)
	for i := 0; i < len(order); i++ {
		order[i] = i // []int{0,1,2,3}
	}
	if s.Variations == 0 {
		// keep the slice
		// log.Printf("shuffler: variations == 0  ->  no shufflings")
	} else {
		class := int64(s.Seed % s.Variations) // user 12; variations 5 => class 2
		src := rand.NewSource(class)          // not ...UTC().UnixNano(), but constant init
		gen := rand.New(src)                  // generator seeded with the class of the user ID

		// This does not cover all elements equally
		if false {
			for i := 0; i < len(order); i++ {
				order[i] = gen.Int() % s.MaxElements // gen.Int 19; max elements 4 =>
				log.Printf("%2v: User %v is class %v => of %v", i, s.Seed, class, order[i])
			}
		}

		swapFct := func(i, j int) {
			order[i], order[j] = order[j], order[i]
		}

		//
		if shufflingRepetitions == 0 {
			shufflingRepetitions = 3 // one repetition would be sufficient as well; just not zero
		}
		shufflingRepetitions = shufflingRepetitions % 7
		for i := 0; i <= shufflingRepetitions; i++ {
			gen.Shuffle(s.MaxElements, swapFct)
			// log.Printf("%2v: User %v is class %v => %+v", i, s.Seed, class, order)
		}
	}

	return order
}
