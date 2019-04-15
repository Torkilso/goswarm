package main

import (
	"fmt"
	"math"
	"sync"
)

var prob *Problem

func main() {
	/*problem := parseFile(2)


	operations := particleSwarmOptimization(problem, 1000, 100)

	drawGannt(problem, operations)
	*/
	prob = parseFile(3)

	target := 56
	targetAcceptable := target + int(0.1*float64(target))

	solutionPSO := runParticleSwarmOptimization(prob)
	solutionBA := BA(targetAcceptable)
	fmt.Println("Makespan; PSO:", solutionPSO.makespan, "BA", solutionBA.makespan)
	visualizeSolutionAsGant(solutionBA.genotype, "gantBA.png", solutionBA.makespan)
	visualizeSolutionAsGant(solutionPSO.genotype, "gantPSO.png", solutionPSO.makespan)
}


func runParticleSwarmOptimization(problem *Problem) Patch {
	n := 6
	solutions := make([]Patch, 0, n)

	channel := make(chan Patch, n)

	var wg sync.WaitGroup
	wg.Add(n * 2)

	for i := 0; i < n; i++ {
		go func(index int) {
			sol := particleSwarmOptimization(problem, 5000, 200)
			fmt.Println("Bes solution from run", sol.makespan)
			channel <- sol
			wg.Done()
		}(i)
	}
	go func() {
		for t := range channel {
			solutions = append(solutions, t)
			wg.Done()
		}
	}()
	wg.Wait()

	bestIdx := 0
	bestVal := math.MaxInt64

	for i := range solutions {
		if solutions[i].makespan < bestVal {
			bestIdx = i
			bestVal = solutions[i].makespan
		}
	}

	return solutions[bestIdx]

}