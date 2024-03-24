package utils

import (
	"math/rand"
	"time"
)

func CatchStatus(baseExperience int) bool {
	// Create local random generated seeded with current time
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	// The max base experience below is a guess
	// It could be higher or lower
	const maxBaseExperience = 300

	// Increate to make harder, 1.0 is baseline
	const difficultyFactor = 1.0

	catchThreshhold := (maxBaseExperience - float64(baseExperience)) / maxBaseExperience * difficultyFactor

	catchProbability := rng.Float64()

	return catchProbability < catchThreshhold
}
