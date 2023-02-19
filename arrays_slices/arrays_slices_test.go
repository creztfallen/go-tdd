package arraysslices

import "testing"

func TestSum(t *testing.T) {

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := Sum(numbers)
		expected := 6

		if result != expected {
			t.Errorf("resultado %d, esperado %d, dado %v", result, expected, numbers)
		}
	})
}
