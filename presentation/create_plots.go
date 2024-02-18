package presentation

import (
	"fmt"

	"github.com/ShugoYoko/playground-tracer/domain"
	"gonum.org/v1/plot/plotter"
)

func Create_polts() plotter.XYs {
	fmt.Print("中心座標を教えてください(例)200.0 200.0>")
	var n [2]float64
	fmt.Scan(&n[0], &n[1]) // コンソールからの入力を取得
	fmt.Print("直線の長さを教えてください(例)80.0>")
	var w float64
	fmt.Scan(&w)
	fmt.Print("回転半径の長さを教えてください(例)40.0>")
	var r float64
	fmt.Scan(&r)

	track := domain.NewTrackInfo(n[0], n[1], w, r)

	points := track.PrintTrackLocation(50)

	return points

}
