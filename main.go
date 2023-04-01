package main

/*
  Brute force approach to calculating approximate probability of rolling at least m matches when rolling 2-14 6-sided dice
*/

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	trials        = 20000 //number of times to roll each handful of dice (t = 1 means roll once - hardly enough data to draw an average from!)
	threshold     = 3     //number of duplicates or more before we say the threshold has been reached
	minDiceToRoll = 2
	maxDiceToRoll = 14
)

var (
	r                *rand.Rand
	thresholdCounter int
)

func main() {

	r = initalizeRandom()

	fmt.Printf("\nData for a threshold value of at least %d matches on handfuls of %d to %d dice\n\n", threshold, minDiceToRoll, maxDiceToRoll)

	//n represents a handful of n dice
	for n := minDiceToRoll; n <= maxDiceToRoll; n++ {

		//run t trials - basically throw n dice t times so we can build a large amount of data
		thresholdCounter = 0
		for t := 1; t <= trials; t++ {

			//roll n dice and store their individual results
			aRoll := make([]int, n)
			for i := 0; i < n; i++ {
				aRoll[i] = d6()
			}

			//look over these results - do we find at least threshold number of duplicates?
			if meetsThreshold(aRoll, threshold) {
				thresholdCounter++
			}
		}

		percent := 100 * (float32(thresholdCounter) / float32(trials))
		fmt.Printf("Rolling %d dice generated %d matches over %d trials or %f%%\n", n, thresholdCounter, trials, percent)
	}

}

func meetsThreshold(aRoll []int, minDupes int) bool {

	//create a counter for the die values 1 - 6 (0 is unused)
	data := make([]int, 7)

	//look through a given roll and count the number of each die face in that roll
	for _, i := range aRoll {

		data[i]++

		//have we seen enough of this face to stop looking?
		if data[i] == minDupes {
			return true
		}
	}

	return false
}

func d6() int {
	return r.Intn(6) + 1
}

func initalizeRandom() *rand.Rand {
	source := rand.NewSource(time.Now().UnixNano())
	return rand.New(source)
}
