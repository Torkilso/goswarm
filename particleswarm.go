package main

import (
	"math"
	"math/rand"
)

func initialize(size int) Swarm {

	particles := make([]*Particle, size)

	for i := range particles {
		particles[i].score = Score{
			position: rand.Float64(),
			velocity: rand.Float64(),
		}
	}
	return particles
}

func evaluateFitness(p *Problem, swarm Swarm) Score {
	globalBest := Score{cost: math.MinInt64, position: 0.0}

	for _, particle := range swarm {
		encoding := decodeGenotype(particle.genotype)
		phenotype := discreteGenoToJobs(p.numJobs, encoding)

		particle.score = Score{
			cost:     cost(p, phenotype),
			position: particle.score.position,
		}

		if globalBest.cost < particle.score.cost {
			globalBest = particle.score
		}
	}
	return globalBest
}

func particleSwarmOptimization(p *Problem, size, iterations int) {

	// Initialize the populaton
	swarm := initialize(size)
	var globalBest Score

	// Do until stopping condition

	for i := 0; i < iterations; i++ {

		// Evaluate fitness and update global best

		globalBest = evaluateFitness(p, swarm)

		for _, particle := range swarm {
			// Update personal best

			if particle.score.cost > particle.personalBest.cost {
				particle.personalBest = particle.score
			}

			// Update velocity
			particle.score.velocity = velocity(particle.score, particle.personalBest, globalBest)

			// Update position
			particle.score.position = particle.score.position + particle.score.velocity

		}
	}
}
