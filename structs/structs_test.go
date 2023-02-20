package structs

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Retângulo", shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "Círculo", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triângulo", shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.shape.Area()
			if result != tt.expected {
				t.Errorf("resultado %.2f, esperado %.2f", result, tt.expected)
			}
		})

	}
}
