package main

import "testing"

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		got := shape.Perimeter()

		if got != expected {
			t.Errorf("%#v got %.2f want %.2f", shape, got, expected)
		}
	}

	perimeterTests := []struct {
		name     string
		shape    Shape
		hasPerim float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 10, Height: 10}, hasPerim: 40.0},
		{name: "circle", shape: Circle{Radius: 10}, hasPerim: 62.83185307179586},
		{name: "triangle", shape: Triangle{Base: 10, Height: 10}, hasPerim: 32.3606797749979},
	}

	for _, tt := range perimeterTests {
		checkPerimeter(t, tt.shape, tt.hasPerim)
	}
}
func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if got != expected {
			t.Errorf("%#v got %.2f want %.2f", shape, got, expected)
		}
	}

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.hasArea)
	}
}
