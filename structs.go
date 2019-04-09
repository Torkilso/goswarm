package main


type Score struct {
	position float64
	velocity float64
	cost 	 int
}
type Particle struct {
	personalBest Score
	score Score
	genotype Genotype

}

type Swarm = []*Particle


type Genotype = []float64

type DiscreteGenotype = []int

type Phenotype = struct {

}


type Problem = struct {
	numJobs, numMachines int
	jobs []Job
}
