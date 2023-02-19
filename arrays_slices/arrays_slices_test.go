package arraysslices

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("resultado %v, esperado %v", result, expected)
	}
}
