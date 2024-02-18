package domain

import (
	"testing"
)

func TestPoint(t *testing.T) {
	expect := Point{X: 3.0, Y: 6.0}

	if expect.X != 3.0 || expect.Y != 6.0 {
		t.Errorf("%f!=3.0", expect.X)
		t.Errorf("%f!=6.0", expect.Y)
	}

}

func TestTrackInfo(t *testing.T) {
	expect := NewTrackInfo(3.0, 4.0, 5.0, 6.0)
	if expect.Point.X != 3.0 || expect.Point.Y != 4.0 {
		t.Errorf("%f!=3.0", expect.Point.X)
		t.Errorf("%f!=4.0", expect.Point.Y)
	}

	if expect.Width != 5.0 || expect.Radius != 6.0 {
		t.Errorf("%f!=5.0", expect.Width)
		t.Errorf("%f!=6.0", expect.Radius)
	}

}
