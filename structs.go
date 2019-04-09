package main


type Score struct {
	position float64
	velocity float64
	cost 	 float64
}
type Particle struct {
	personalBest Score
	score Score
	genotype Genotype

}

type Swarm = struct {
	particles []*Particle
	globalBest Score
}

type Genotype = []float64

type DiscreteGenotype = []int

type Phenotype = struct {

}

