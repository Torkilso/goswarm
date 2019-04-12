package main

import (
	"github.com/fogleman/gg"
	"image/color"
	"strconv"
)

type Entry struct {
	machine   int
	job       int
	operation int
	start     int
	end       int
}

var (
	imageWidth = 2100
)

func drawEntry(dc *gg.Context, entry *Entry, makespan int) {
	x := float64(1 + ((2000)*entry.start)/makespan)
	y := float64(11 + entry.machine*100)
	width := float64(((2000)*entry.end)/makespan) - x

	dc.DrawRectangle(51+x, y, width, 98)
	dc.SetColor(color.RGBA{R: 0, G: 0, B: uint8((float64(entry.job) / float64(prob.numJobs)) * 255), A: 255})
	dc.Fill()

	dc.SetColor(color.White)
	text := strconv.Itoa(entry.operation) + "/" + strconv.Itoa(entry.job)
	dc.DrawStringAnchored(text, 51+x+width/2, y+50, 0.5, 0.5)
	dc.Stroke()
}

func fillGant(dc *gg.Context, genotype Genotype, makespan int) {

	discrete := decodeGenotype(genotype)
	operations := discreteGenoToJobs(prob.numJobs, discrete)

	entries := make([]Entry, 0)

	machinesIdleAt := make(map[int]int)
	jobsSchedules := make(map[int]*JobSchedule, prob.numJobs)

	for i := 0; i < prob.numMachines; i++ {
		machinesIdleAt[i] = 0
	}

	for i := 0; i < prob.numJobs; i++ {
		jobsSchedules[i] = &JobSchedule{0, 0}
	}

	for _, operation := range operations {
		operationForJob := jobsSchedules[operation].operationsFinished
		lastOperationFinishedAt := jobsSchedules[operation].lastOperationFinishedAt
		workload := prob.jobs[operation][operationForJob]
		machineIdleAt := machinesIdleAt[workload.machine]

		entry := Entry{workload.machine, operation, operationForJob, 0, 0}

		if machineIdleAt <= lastOperationFinishedAt {

			entry.start = lastOperationFinishedAt
			entry.end = lastOperationFinishedAt + workload.duration

			machinesIdleAt[workload.machine] = lastOperationFinishedAt + workload.duration
			jobsSchedules[operation].lastOperationFinishedAt += workload.duration

		} else {

			entry.start = machineIdleAt
			entry.end = machineIdleAt + workload.duration

			machinesIdleAt[workload.machine] += workload.duration
			jobsSchedules[operation].lastOperationFinishedAt = machineIdleAt + workload.duration
		}

		entries = append(entries, entry)

		jobsSchedules[operation].operationsFinished += 1
	}

	for _, entry := range entries {
		drawEntry(dc, &entry, makespan)
	}
}

func visualizeSolutionAsGant(genotype Genotype, filename string, makespan int) {

	imageHeight := prob.numMachines*100 + 50

	dc := gg.NewContext(imageWidth, imageHeight)

	dc.DrawRectangle(0, 0, float64(imageWidth), float64(imageHeight))
	dc.SetColor(color.White)
	dc.Fill()

	// Draw vertical dashed lines
	/*for i := 0; i < 100; i++ {
		dc.DrawLine(float64(50+i*20), float64(10), float64(50+i*20), float64(imageHeight-40))
	}*/

	dc.SetColor(color.Black)
	dc.SetDash(4)
	dc.SetLineWidth(1)
	dc.Stroke()

	// Draw vertical lines and time
	for i := 0; i <= 20; i++ {
		time := (i * 5 * makespan) / 100
		dc.DrawStringAnchored(strconv.Itoa(time), float64(50+i*100), float64(imageHeight-20), 0.5, 0.5)
		dc.DrawLine(float64(50+i*100), float64(10), float64(50+i*100), float64(imageHeight-40))
	}

	dc.SetColor(color.Black)
	dc.SetDash()
	dc.SetLineWidth(1)
	dc.Stroke()

	// Draw horizontal lines and machine numbers
	for i := 0; i < prob.numMachines+1; i++ {
		dc.DrawStringAnchored("M"+strconv.Itoa(i), 25, float64(55+i*100), 0.5, 0.5)
		dc.DrawLine(float64(50), float64(10+i*100), float64(imageWidth-50), float64(10+i*100))
	}

	dc.SetColor(color.Black)
	dc.SetLineWidth(1)
	dc.Stroke()

	fillGant(dc, genotype, makespan)

	dc.SavePNG(filename)
}
