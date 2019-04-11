package main

import (
	"math"
	"math/rand"
	"sort"
)

type GenoDecoder struct {
	index int
	value float64
}

func decodeGenotype(geno Genotype) (encoding DiscreteGenotype) {
	genoDecoder := make([]GenoDecoder, len(geno))

	for i, val := range geno {
		genoDecoder[i] = GenoDecoder{index: i, value: val}
	}

	sort.Slice(genoDecoder, func(i, j int) bool {
		return genoDecoder[i].value < genoDecoder[j].value
	})

	encoding = make(DiscreteGenotype, len(geno))

	for i, decoder := range genoDecoder {
		encoding[decoder.index] = i
	}
	return encoding
}
func discreteGenoToJobs(numJobs int, encoding DiscreteGenotype) []int {
	jobs := make([]int, len(encoding))

	for i, val := range encoding {
		jobs[i] = val % numJobs
	}

	return jobs
}

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
