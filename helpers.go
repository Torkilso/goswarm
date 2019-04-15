package main

import (
	"math"
	"math/rand"
	"strconv"
)

var  (
	maxVelocity float64 = 10
	minVelocity float64 = -10
	selfLearning float64 = 10
	socialLearning float64 = 10
	inertia float64 = 10

)


type Operation = struct {
	machine int
	start int
	stop int
	id string
	job int
	numberInLine int
}


func phenoToOperations(p *Problem, jobs Phenotype) []*Operation{

	operations := make([]*Operation, p.numJobs * p.numMachines)
	machineDurations := make([]int, p.numMachines)
	jobCheckpoints := make([]int, p.numMachines)

	operationNumbers := make([]int, p.numJobs)
	for i, job := range jobs {
		operationNum := operationNumbers[job]
		operation := p.jobs[job][operationNum]

		start := int(math.Max(float64(machineDurations[operation.machine]), float64(jobCheckpoints[job])))
		operations[i] = &Operation{
			machine: operation.machine,
			start:   start,
			stop:    start + operation.duration,
			id: "Job: " + strconv.Itoa(job) + ", Op: " + strconv.Itoa(operationNum),
			job: job,
			numberInLine: operationNumbers[job],
		}
		machineDurations[operation.machine] = operations[i].stop
		jobCheckpoints[job] = operations[i].stop

		operationNumbers[job] += 1
	}
	return operations

}

func cost(operations []*Operation) int {
	max := math.MinInt64

	for _, operation := range operations {
		if operation.stop > max {
			max = operation.stop
		}
	}
	return max

}

func velocity(last *Particle, globalBest *Particle) []float64 {

	personalDiff := sub(last.personalBest.genotype, last.genotype)
	globalDiff := sub(globalBest.genotype, last.genotype)

	velocityFromLast := mult(last.velocity, inertia)
	velocityFromPersonal := mult(personalDiff, selfLearning * rand.Float64())
	velocityFromGlobal := mult(globalDiff, socialLearning * rand.Float64())

	newVelocity :=  add(velocityFromLast, add(velocityFromGlobal, velocityFromPersonal))

	newVelocity = max(newVelocity, maxVelocity)
	newVelocity = min(newVelocity, minVelocity)

	return newVelocity
}



func add(this, other []float64) []float64 {
	newArr := make([]float64, len(this))
	for i := range this {
		newArr[i] = this[i] + other[i]
	}
	return newArr

}
func sub(this, other []float64) []float64 {
	newArr := make([]float64, len(this))
	for i := range this {
		newArr[i] = this[i] - other[i]
	}
	return newArr

}
func mult(this []float64, mult float64) []float64 {
	newArr := make([]float64, len(this))
	for i := range this {
		newArr[i] = this[i] * mult
	}
	return newArr

}

func max(this []float64, value float64) []float64 {
	newArr := make([]float64, len(this))
	for i := range this {
		newArr[i] = math.Max(this[i], value)
	}
	return newArr
}

func min(this []float64, value float64) []float64 {
	newArr := make([]float64, len(this))
	for i := range this {
		newArr[i] = math.Min(this[i], value)
	}
	return newArr
}
