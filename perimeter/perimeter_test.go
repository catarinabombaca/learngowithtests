package perimeter

import (
	"github.com/catarinabombaca/learngowithtests/testhelpers"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	testhelpers.AssertCorrectFloat(t, got, want)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: &Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72},
		{name: "Circle", shape: &Circle{radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: &Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			testhelpers.AssertCorrectFloat(t, got, test.hasArea)
		})
	}
}
