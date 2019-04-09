package main

import (
	"math/rand"
)

var  (
	maxVelocity float64 = 10
	minVelocity float64 = 10
	selfLearning float64 = 10
	socialLearning float64 = 10
	inertia float64 = 10

)
func cost(pheno *Phenotype) float64 {

}

func velocity(last, personalBest, globalBest Score) float64 {
	newVelocity := inertia * last.velocity +
						selfLearning * rand.Float64() * (personalBest.position - last.position) +
						socialLearning * rand.Float64() * (globalBest.position - last.position)

	if newVelocity > maxVelocity {
		newVelocity = maxVelocity
	}else if newVelocity < minVelocity {
		newVelocity = minVelocity
	}

	return newVelocity
}