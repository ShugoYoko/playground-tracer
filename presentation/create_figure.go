package presentation

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Create_figure(points plotter.XYs) {

	p := plot.New()

	p.Title.Text = "グラウンド"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"グラウンド", points,
	)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}

}
