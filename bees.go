package main

import (
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
	bestPatchesAmount      = 10  // m
	elitePatches           = 5   // e
	beesForElitePatches    = 5   // nep
	beesForNonElitePathces = 5   // nsp
	neighbourHoodSize      = 5   // ngh
	beesGenerations        = 100
)

func (p *Patch) evaluateFitnessValue() {

	p.makespan = 1
}

func (p *Patches) evaluateFitnessValue() {
	for _, patch := range *p {
		patch.evaluateFitnessValue()
	}
}

func generatePatch(jobs, machines int) *Patch {

	genotype := make([]float64, jobs*machines)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := range genotype {
		genotype[i] = r1.Float64() * 100
	}

	return &Patch{
		makespan: 0,
		genotype: genotype,
	}
}

func generatePatches(amount, jobs, machines int) Patches {
	patches := make([]*Patch, amount)

	for i := range patches {
		patch := generatePatch(jobs, machines)
		patches[i] = patch
	}

	return patches
}

func (p *Patches) best() Patches {
	return (*p)[:bestPatchesAmount]
}

func (p *Patches) nonBest() Patches {
	return (*p)[bestPatchesAmount:]
}

func (p *Patches) elites() Patches {
	return (*p)[:elitePatches]
}

func (p *Patches) nonElites() Patches {
	return (*p)[elitePatches:]
}

func (p *Patch) neighbours() Patches {
	return make([]*Patch, 0)
}

func (p *Patches) neighbours() Patches {
	return *p

	// TODO create neighbourhood from each patch
}

func BA(problem *Problem) Patch {
	patches := generatePatches(beesAmount, problem.numJobs, problem.numMachines)
	patches.evaluateFitnessValue()

	sort.Slice(patches, func(i, j int) bool {
		return (patches)[i].makespan < (patches)[j].makespan
	})

	for i := 0; i <= beesGenerations; i++ {
		bestPatches := patches.best()

		elitePatches := bestPatches.elites()
		nonElitePatches := bestPatches.nonElites()

		foragerElitePatches := elitePatches.neighbours()
		foragerNonElitePatches := nonElitePatches.neighbours()

		newPatches := generatePatches(beesAmount-beesForElitePatches-beesForNonElitePathces, problem.numJobs, problem.numMachines)

		patches = append(foragerElitePatches, foragerNonElitePatches...)
		patches = append(patches, newPatches...)

		patches.evaluateFitnessValue()

		sort.Slice(patches, func(i, j int) bool {
			return (patches)[i].makespan < (patches)[j].makespan
		})
	}

	return *patches[0]
}
