package domain

import (
	"math"

	"gonum.org/v1/plot/plotter"
)

type Point struct {
	X float64
	Y float64
}

// TrackInfo 構造体の定義
type TrackInfo struct {
	Point
	Width  float64
	Radius float64
}

// NewTrackInfo 関数の定義
func NewTrackInfo(x, y, width, radius float64) *TrackInfo {
	return &TrackInfo{
		Point:  Point{X: x, Y: y},
		Width:  width,
		Radius: radius,
	}
}

func (ti *TrackInfo) PrintTrackLocation(sp int) plotter.XYs {
	fsp := float64(sp)

	points := make(plotter.XYs, (sp+1)*4)
	total_count := 0
	//上軸横(<-)
	ux := ti.Point.X + 1.0/2.0*ti.Width

	for i := 0; i <= sp; i++ {
		points[total_count].X = ux - ti.Width/fsp*float64(i)
		points[total_count].Y = ti.Point.Y + ti.Radius
		total_count = total_count + 1
	}

	//左カーブ
	lcx := ti.Point.X - 1.0/2.0*ti.Width
	lcy := ti.Point.Y
	langle := 90.0

	for i := 0; i <= sp; i++ {
		fi := float64(i)
		points[total_count].X = lcx + ti.Radius*math.Cos((langle+180/fsp*fi)*(math.Pi/180))
		points[total_count].Y = lcy + ti.Radius*math.Sin((langle+180/fsp*fi)*(math.Pi/180))
		total_count = total_count + 1
	}

	//下軸(->)
	lx := ti.Point.X - 1.0/2.0*ti.Width

	for i := 0; i <= sp; i++ {
		points[total_count].X = lx + ti.Width/fsp*float64(i)
		points[total_count].Y = ti.Point.Y - ti.Radius
		total_count = total_count + 1
	}

	//右カーブ
	rcx := ti.Point.X + 1.0/2.0*ti.Width
	rcy := ti.Point.Y
	rangle := 270.0

	for i := 0; i <= sp; i++ {
		fi := float64(i)
		points[total_count].X = rcx + ti.Radius*math.Cos((rangle+180/fsp*fi)*(math.Pi/180))
		points[total_count].Y = rcy + ti.Radius*math.Sin((rangle+180/fsp*fi)*(math.Pi/180))
		total_count = total_count + 1
	}

	return points
}
