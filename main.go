package main

import (
	"fmt"

	"github.com/ShugoYoko/playground-tracer/presentation"
)

func main() {

	points := presentation.Create_polts()

	fmt.Print("座標出力:1,グラフ出力:2>")
	var n int
	fmt.Scan(&n)

	if n == 1 {
		presentation.Print_polts(points)
	} else if n == 2 {
		presentation.Create_figure(points)
	} else {
		fmt.Println("1か2を選択してください")
	}
}
