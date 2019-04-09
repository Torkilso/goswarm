package main

import (
	"math"
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

func encodingToPhenotype(encoding DiscreteGenotype) (pheno Phenotype) {

}


func initialize(size int) *Swarm{


	return nil
}


func evaluateFitness(swarm *Swarm ) Score {
	globalBest := Score{cost: -math.MaxFloat64, position: 0.0}

	for _, particle := range swarm.particles {
		encoding := decodeGenotype(particle.genotype)
		phenotype := encodingToPhenotype(encoding)
		particle.score = Score{cost: cost(&phenotype), position: particle.score.position}

		if globalBest.cost < particle.score.cost {
			globalBest = particle.score
		}
	}
	return globalBest
}



func particleSwarmOptimization(size, iterations int){

	// Initialize the populaton
	swarm := initialize(size)
	var globalBest Score

	// Do until stopping condition

	for i := 0 ; i < iterations ; i++ {

		// Evaluate fitness and update global best

		globalBest = evaluateFitness(swarm)

		for _, particle := range swarm.particles {
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