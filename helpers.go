package main

import (
	"github.com/google/gxui/math"
	"math/rand"
)

var (
	maxVelocity    float64 = 10
	minVelocity    float64 = -10
	selfLearning   float64 = 10
	socialLearning float64 = 10
	inertia        float64 = 10
)

func cost(p *Problem, jobs []int) int {
	machineDurations := make([]int, p.numMachines)

	for _, job := range jobs { // TODO: This is not working... (makes no sense)
		machineDurations[job] += p.jobs[job][0].duration
	}

	max := math.MinInt

	for _, dur := range machineDurations {
		if dur > max {
			max = dur
		}
	}
	return max

}

func velocity(last, personalBest, globalBest Score) float64 {
	newVelocity := inertia*last.velocity +
		selfLearning*rand.Float64()*(personalBest.position-last.position) +
		socialLearning*rand.Float64()*(globalBest.position-last.position)

	if newVelocity > maxVelocity {
		newVelocity = maxVelocity
	} else if newVelocity < minVelocity {
		newVelocity = minVelocity
	}

	return newVelocity
}
