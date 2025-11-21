// this isnt a real test file
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetClipOrder(numOfClips int) []int {
	// Create a new random generator using the provided seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create the slice [1..numOfClips]
	newClipOrder := make([]int, numOfClips)
	for i := 0; i < numOfClips; i++ {
		newClipOrder[i] = i + 1
	}

	fmt.Println("orginal file order: ", newClipOrder)

	// Shuffle deterministically based on the seed
	r.Shuffle(len(newClipOrder), func(i, j int) {
		newClipOrder[i], newClipOrder[j] = newClipOrder[j], newClipOrder[i]
	})

	return newClipOrder
}

func main() {

	fmt.Println("new file order: ", GetClipOrder(14))
}
