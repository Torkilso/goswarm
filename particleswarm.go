package main

import (
	"fmt"
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


func initialize(p *Problem, size int) Swarm{

	particles := make([]*Particle, size)


	for i := range particles {
		geno := make(Genotype, p.numMachines * p.numJobs)

		for i := range geno {
			geno[i] = rand.Float64()
		}
		particles[i] = &Particle{
			genotype: geno,
		}
	}
	return particles
}

func genoToPheno(p *Problem, geno Genotype) Phenotype {
	encoding := decodeGenotype(geno)
	return discreteGenoToJobs(p.numJobs, encoding)

}


func evaluateFitness(p *Problem, swarm Swarm) *Particle {
	globalBest := &Particle{cost: math.MaxInt64}

	for _, particle := range swarm {
		phenotype := genoToPheno(p, particle.genotype)

		particle.cost = cost(phenoToOperations(p, phenotype))


		if globalBest.cost > particle.cost {
			globalBest = particle
		}
	}
	return globalBest
}



func particleSwarmOptimization(p *Problem, size, iterations int) []*Operation {

	// Initialize the populaton
	swarm := initialize(p, size)
	var globalBest *Particle

	// Do until stopping condition

	for i := 0 ; i < iterations ; i++ {

		// Evaluate fitness and update global best

		globalBest = evaluateFitness(p, swarm)

		for _, particle := range swarm {
			// Update personal best

			if particle.cost < particle.personalBest.cost {
				particle.personalBest = *particle
			}

			// Update velocity
			particle.velocity = velocity(particle, globalBest)

			// Update position
			particle.genotype = add(particle.genotype, particle.velocity)

			}
		fmt.Println("Global best", globalBest.cost)

	}
	// Find best

	best := math.MaxInt64
	bestIdx := 0
	for i, particle := range swarm {
		if particle.cost < best {
			best = particle.cost
			bestIdx = i
		}
	}
	return phenoToOperations(p, genoToPheno(p, swarm[bestIdx].genotype))
}