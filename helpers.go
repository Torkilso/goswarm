package main

import (
	"math"
	"math/rand"
	"sort"
)

var (
	maxVelocity    float64 = 10
	minVelocity    float64 = -10
	selfLearning   		   = 2.0
	socialLearning 		   = 2.0
	startInertia		   = 0.9
	endInertia			   = 0.4
)

type Operation = struct {
	machine int
	start int
	stop int
	id string
	job int
	numberInLine int
}


func phenoToOperations(p *Problem, jobs Phenotype) []*Operation {
	return make([]*Operation, 0)
}
type JobSchedule struct {
	lastOperationFinishedAt int
	operationsFinished      int
}

func (s *OperationSequence) makespan() int {

	machinesIdleAt := make(map[int]int)
	jobsSchedules := make(map[int]*JobSchedule, prob.numJobs)

	for i := 0; i < prob.numMachines; i++ {
		machinesIdleAt[i] = 0
	}

	for i := 0; i < prob.numJobs; i++ {
		jobsSchedules[i] = &JobSchedule{0, 0}
	}

	for _, operation := range *s {

		operationForJob := jobsSchedules[operation].operationsFinished
		lastOperationFinishedAt := jobsSchedules[operation].lastOperationFinishedAt

		workload := prob.jobs[operation][operationForJob]
		machineIdleAt := machinesIdleAt[workload.machine]

		if machineIdleAt <= lastOperationFinishedAt {
			machinesIdleAt[workload.machine] = lastOperationFinishedAt + workload.duration
			jobsSchedules[operation].lastOperationFinishedAt += workload.duration
		} else {
			machinesIdleAt[workload.machine] += workload.duration
			jobsSchedules[operation].lastOperationFinishedAt = machineIdleAt + workload.duration
		}

		jobsSchedules[operation].operationsFinished += 1
	}

	maxDuration := 0
	for _, job := range jobsSchedules {
		if job.lastOperationFinishedAt > maxDuration {
			maxDuration = job.lastOperationFinishedAt
		}
	}

	return maxDuration
}

func velocity(last *Particle, globalBest *Particle, percentDone float64) []float64 {

	personalDiff := sub(last.personalBest.genotype, last.genotype)
	globalDiff := sub(globalBest.genotype, last.genotype)

	inertia := ((startInertia - endInertia) * percentDone) + endInertia

	velocityFromLast := mult(last.velocity, inertia)
	velocityFromPersonal := mult(personalDiff, selfLearning * rand.Float64())
	velocityFromGlobal := mult(globalDiff, socialLearning * rand.Float64())


	newVelocity :=  add(velocityFromLast, add(velocityFromGlobal, velocityFromPersonal))

	newVelocity = max(newVelocity, minVelocity)
	newVelocity = min(newVelocity, maxVelocity)

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

func discreteGenoToJobs(numJobs int, encoding DiscreteGenotype) OperationSequence {
	jobs := make([]int, len(encoding))

	for i, val := range encoding {
		jobs[i] = val % numJobs
	}

	return jobs
}

