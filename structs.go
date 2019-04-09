package main


type Score struct {
	cost float64
	position float64
}
type Particle struct {
	personalBest Score
	score Score

}

type Swarm = struct {
	particles []*Particle
	globalBest Score
}

type Genotypes = [][]float64