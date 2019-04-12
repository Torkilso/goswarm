package main

var prob *Problem

func main() {
	prob = parseFile(2)

	//particleSwarmOptimization(problem, 100, 100)
	solution := BA()
	visualizeSolutionAsGant(solution.genotype, "gant.png", solution.makespan)
}
