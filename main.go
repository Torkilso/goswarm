package main

import "fmt"

var prob *Problem

func main() {
	prob = parseFile(1)

	target := 56
	targetAcceptable := target + int(0.1*float64(target))

	//particleSwarmOptimization(problem, 100, 100)
	solution := BA(targetAcceptable)
	fmt.Println("makespan:", solution.makespan)
	visualizeSolutionAsGant(solution.genotype, "gant.png", solution.makespan)
}
