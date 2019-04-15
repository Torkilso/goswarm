package main

import (
	"fmt"
	"math"
	"math/rand"
)


func initialize(p *Problem, size int) Swarm{
	particles := make([]*Particle, size)


	for i := range particles {
		geno := make(Genotype, p.numMachines * p.numJobs)
		velocity := make([]float64, len(geno))
		for i := range geno {
			geno[i] = rand.Float64()
			velocity[i] = rand.Float64()
		}
		particles[i] = &Particle{
			genotype: geno,
			velocity: velocity,
		}
	}
	return particles
}

func (p *Particle) evaluateFitnessValue() int {
	decoded := decodeGenotype(p.genotype)
	operationSequence := discreteGenoToJobs(prob.numJobs, decoded)
	p.cost = operationSequence.makespan()

	return p.cost
}


func evaluateFitness(swarm Swarm) *Particle {
	globalBest := &Particle{cost: math.MaxInt64}

	for _, particle := range swarm {

		particle.evaluateFitnessValue()

		if globalBest.cost > particle.cost {
			globalBest = particle
		}
	}
	return globalBest
}



func particleSwarmOptimization(p *Problem, size, iterations int) Patch {

	// Initialize the populaton
	swarm := initialize(p, size)
	var globalBest *Particle

	// Do until stopping condition

	for i := 0; i < iterations; i++ {

		// Evaluate fitness and update global best

		thisGlobalBest := evaluateFitness(swarm)
		if globalBest == nil {
			globalBest = thisGlobalBest
		}else if thisGlobalBest.cost < globalBest.cost {
			genoCopy := make(Genotype, len(thisGlobalBest.genotype))
			copy(genoCopy, thisGlobalBest.genotype)
			globalBest = &Particle{cost: thisGlobalBest.cost, genotype: genoCopy}
		}

		for _, particle := range swarm {
			// Update personal best
			if particle.personalBest == nil {
				particle.personalBest = particle
			} else if particle.cost < particle.personalBest.cost {
				particle.personalBest = particle
			}

			// Update velocity
			particle.velocity = velocity(particle, globalBest, 1.0 - float64(i) / float64(iterations))

			// Update position
			particle.genotype = add(particle.genotype, particle.velocity)

			}
		if i % 50 == 0 {
			fmt.Println("Global best", globalBest.cost)
		}

	}

	return Patch{makespan: globalBest.cost, genotype: globalBest.genotype}



}
