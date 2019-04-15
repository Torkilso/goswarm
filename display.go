package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

func drawGannt(pr *Problem, operations []*Operation) {

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "JSSP"
	p.X.Label.Text = "Time units"
	p.Y.Label.Text = "Machines"
	// Draw a grid behind the data
	p.Add(plotter.NewGrid())

	lines := make([]*plotter.Line, len(operations))



	for i, op := range operations {
		line := make(plotter.XYs, 2)
		line[0].Y = float64(op.machine)
		line[1].Y = float64(op.machine)
		line[0].X = float64(op.start)
		line[1].X = float64(op.stop)
		lines[i], _ = plotter.NewLine(line)

		c := 0xFF - uint8(0xFF * op.job / pr.numJobs)
		lines[i].Color = color.RGBA{R: c ,G: 0, B: 0,A: 0xFF}
		lines[i].Width = 30
	}


	for _, line := range lines {
		p.Add(line)
	}

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 8*vg.Inch, "points.png"); err != nil {
		panic(err)
	}


}



