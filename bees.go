package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
i = 0
Generate initial population.
Evaluate Fitness Value of initial population.
Sort the initial population based on the fitness result.
While i <= MaxIter || FitnessValue_i - FitnessValue_i-1 <= Error:
	1. i = i + l;
	2. Select the elite patches and non-elite best patches for neighborhood search.
	3. Recruit the forager bees to the elite patches and non-elite best patches.
	4. Evaluate the fitness value of each patch.
	5. Sort the results based on their fitness.
	6. Allocate the rest of the bees for global search to the non-best locations.
	7. Evaluate the fitness value of non-best patches.
	8. Sort the overall results based on their fitness.
	9. Run the algorithm until termination criteria met.
End

Constants:
Number of scout bees in the selected patches n
Number of best patches in the selected patches m
Number of elite patches in the selected best patches e
Number of recruited bees in the elite patches nep
Number of recruited bees in the non-elite best patches nsp
The size of neighborhood for each patch ngh
Number of iterations Maxiter
Difference between value of the first and last iterations diff
*/

var (
	beesAmount             = 100 // n
	bestPatchesAmount      = 15  // m
	elitePatchesAmount     = 3   // e
	beesForElitePatches    = 20  // nep
	beesForNonElitePatches = 5   // nsp
	neighbourHoodSize      = 5   // ngh
	beesGenerations        = 1000
)

func (p *Patch) evaluateFitnessValue() int {
	decoded := decodeGenotype(p.genotype)
	operationSequence := discreteGenoToJobs(prob.numJobs, decoded)
	p.makespan = operationSequence.makespan()

	return p.makespan
}

func (p *Patches) evaluateFitnessValue() {
	for _, patch := range *p {
		patch.evaluateFitnessValue()
	}
}

func generatePatch(jobs, machines, seed int) *Patch {

	genotype := make([]float64, jobs*machines)

	s1 := rand.NewSource(time.Now().UnixNano() + int64(seed))
	r1 := rand.New(s1)

	for i := range genotype {
		genotype[i] = r1.Float64() * 100
	}

	return &Patch{
		makespan: 0,
		genotype: genotype,
	}
}

func generatePatches(amount int) Patches {
	patches := make([]*Patch, amount)

	for i := range patches {
		patch := generatePatch(prob.numJobs, prob.numMachines, i)
		patches[i] = patch
	}

	return patches
}

func (p *Patch) swapMutation(indexA, indexB int) {
	p.genotype[indexA], p.genotype[indexB] = p.genotype[indexB], p.genotype[indexA]
}

func (p *Patch) insertMutation(from, to int) {

}

func (p *Patch) inversionMutation(index, length int) {

}

func (p *Patch) longDistanceMutation(from, to, length int) {

}

func (p *Patch) neighbour(seed int) *Patch {
	s1 := rand.NewSource(time.Now().UnixNano() + int64(seed))
	r1 := rand.New(s1)

	neighbour := Patch{}
	neighbour.genotype = append([]float64(nil), p.genotype...)

	mutationType := r1.Intn(1)

	// only swap is used at the moment
	if mutationType == 0 {
		indexA := r1.Intn(len(neighbour.genotype) - 1)
		indexB := r1.Intn(len(neighbour.genotype) - 1)
		neighbour.swapMutation(indexA, indexB)
	} else if mutationType == 1 {
		from := r1.Intn(len(neighbour.genotype) - 1)
		to := r1.Intn(len(neighbour.genotype) - 1)
		neighbour.insertMutation(from, to)
	} else if mutationType == 2 {
		index := r1.Intn(len(neighbour.genotype) - 1)
		length := r1.Intn(len(neighbour.genotype) - 1)
		neighbour.inversionMutation(index, length)
	} else if mutationType == 3 {
		from := r1.Intn(len(neighbour.genotype) - 1)
		to := r1.Intn(len(neighbour.genotype) - 1)
		length := r1.Intn(len(neighbour.genotype) - 1)
		neighbour.longDistanceMutation(from, to, length)
	}

	neighbour.evaluateFitnessValue()
	return &neighbour
}

func (p *Patch) neighbours(size int) Patches {
	neighbours := make([]*Patch, size)

	neighbours[0] = p

	for i := 1; i < size; i++ {
		neighbours[i] = p.neighbour(i)
		//fmt.Println("makespan", neighbours[i].makespan)
	}

	return neighbours
}

func (p *Patches) neighbours(size int) Patches {
	neighbours := make([]*Patch, 0)

	//fmt.Println("finding neighbours")

	for _, patch := range *p {
		patchNeighbours := patch.neighbours(size)

		sort.Slice(patchNeighbours, func(i, j int) bool {
			return (patchNeighbours)[i].makespan < (patchNeighbours)[j].makespan
		})

		neighbours = append(neighbours, patchNeighbours[0])
	}

	return neighbours[:len(*p)]
}

func BA(target int) Patch {
	patches := generatePatches(beesAmount)
	patches.evaluateFitnessValue()

	sort.Slice(patches, func(i, j int) bool {
		return (patches)[i].makespan < (patches)[j].makespan
	})

	for i := 0; i <= beesGenerations; i++ {

		bestPatches := patches[:bestPatchesAmount]

		elitePatches := bestPatches[:elitePatchesAmount]
		nonElitePatches := bestPatches[elitePatchesAmount:]

		foragerElitePatches := elitePatches.neighbours(beesForElitePatches)
		foragerNonElitePatches := nonElitePatches.neighbours(beesForNonElitePatches)

		newPatches := generatePatches(beesAmount - bestPatchesAmount)

		patches = append(foragerElitePatches, foragerNonElitePatches...)
		patches = append(patches, newPatches...)

		patches.evaluateFitnessValue()

		sort.Slice(patches, func(i, j int) bool {
			return (patches)[i].makespan < (patches)[j].makespan
		})

		fmt.Printf("\rGeneration %d, best makespan: %d", i, patches[0].makespan)

		// uncomment to exit on target condition
		/*if patches[0].makespan < target {
			fmt.Println("\nFound solution below target")
			return *patches[0]
		}*/
	}

	return *patches[0]
}
