package presentation

import (
	"fmt"

	"gonum.org/v1/plot/plotter"
)

func Print_polts(plots plotter.XYs) {
	for _, v := range plots {
		fmt.Println("(", v.X, ",", v.Y, ")")
	}
	//標準入力待ち
	var stop float64
	fmt.Scan(&stop)

}
