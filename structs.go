package main


type Particle struct {
	personalBest Particle
	velocity []float64
	cost 	 int
	genotype Genotype

}

type Swarm = []*Particle


type Genotype = []float64

type DiscreteGenotype = []int

type Phenotype = []int


type Problem = struct {
	numJobs, numMachines int
	jobs []Job
}
