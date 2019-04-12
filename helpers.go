package main

import (
	"github.com/google/gxui/math"
	"math/rand"
	"sort"
)

var (
	maxVelocity    float64 = 10
	minVelocity    float64 = -10
	selfLearning   float64 = 10
	socialLearning float64 = 10
	inertia        float64 = 10
)

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

func cost(p *Problem, jobs []int) int {
	machineDurations := make([]int, p.numMachines)

	for _, job := range jobs { // TODO: This is not working... (makes no sense)
		machineDurations[job] += p.jobs[job][0].duration
	}

	max := math.MinInt

	for _, dur := range machineDurations {
		if dur > max {
			max = dur
		}
	}
	return max

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

func velocity(last, personalBest, globalBest Score) float64 {
	newVelocity := inertia*last.velocity +
		selfLearning*rand.Float64()*(personalBest.position-last.position) +
		socialLearning*rand.Float64()*(globalBest.position-last.position)

	if newVelocity > maxVelocity {
		newVelocity = maxVelocity
	} else if newVelocity < minVelocity {
		newVelocity = minVelocity
	}

	return newVelocity
}
