package main

import "log"

func main() {
	numMachines, jobs := parseFile(1)

	log.Println(numMachines, jobs)


	particleSwarmOptimization()
}