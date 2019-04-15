package main

type Particle struct {
	personalBest *Particle
	velocity     []float64
	cost         int
	genotype     Genotype
}

type Score struct {
	position float64
	velocity float64
	cost     int
}



type Swarm = []*Particle
type Genotype = []float64
type DiscreteGenotype = []int

type Phenotype = []int


type Patch struct {
	makespan int
	genotype Genotype
}

type Patches []*Patch

type OperationSequence []int

type Problem = struct {
	numJobs, numMachines int
	jobs                 []Job
}
