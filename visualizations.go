package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"strconv"
)

func visualizeSolutionAsGant(problem *Problem, genotype Genotype, filename string) {

	discrete := decodeGenotype(genotype)
	operations := discreteGenoToJobs(problem.numJobs, discrete)

	imageWidth := 1500
	imageHeight := problem.numMachines*100 + 50

	dc := gg.NewContext(imageWidth, imageHeight)

	dc.DrawRectangle(0, 0, float64(imageWidth), float64(imageHeight))
	dc.SetColor(color.White)
	dc.Fill()

	// Draw vertical dashed lines
	for i := 0; i < 100; i++ {
		dc.DrawLine(float64(50+i*20), float64(10), float64(50+i*20), float64(imageHeight-40))
	}

	dc.SetColor(color.Black)
	dc.SetDash(6)
	dc.SetLineWidth(1)
	dc.Stroke()

	// Draw vertical lines and time
	for i := 0; i < 15; i++ {
		dc.DrawStringAnchored(strconv.Itoa(i*5), float64(50+i*100), float64(imageHeight-20), 0.5, 0.5)
		dc.DrawLine(float64(50+i*100), float64(10), float64(50+i*100), float64(imageHeight-40))
	}

	dc.SetColor(color.Black)
	dc.SetDash()
	dc.SetLineWidth(1)
	dc.Stroke()

	// Draw horizontal lines and machine numbers
	for i := 0; i < problem.numMachines+1; i++ {
		dc.DrawStringAnchored("M"+strconv.Itoa(i), 25, float64(55+i*100), 0.5, 0.5)
		dc.DrawLine(float64(50), float64(10+i*100), float64(imageWidth), float64(10+i*100))
	}

	dc.SetColor(color.Black)
	dc.SetLineWidth(1)
	dc.Stroke()

	for _, op := range operations {
		fmt.Print(op, ": ")

		for _, wl := range problem.jobs[op] {

			fmt.Print(wl.machine, " ", wl.duration, " ")
		}

		fmt.Println()
	}

	dc.SavePNG(filename)
}
