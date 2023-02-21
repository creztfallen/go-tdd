package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"teste": "isso é apenas um teste"}

	t.Run("palavra conhecida", func(t *testing.T) {
		result, _ := dictionary.Search("teste")
		expected := "isso é apenas um teste"

		compareStrings(t, result, expected)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, result := dictionary.Search("desconhecida")

		compareError(t, result, errWordNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("palavra existente", func(t *testing.T) {
		key := "teste"
		value := "isso é apenas um teste"
		newValue := "nova definição"
		dictionary := Dictionary{key: value}
		err := dictionary.Update(key, newValue)

		compareError(t, err, nil)
		compareValue(t, dictionary, key, newValue)
	})
	t.Run("palavra nova", func(t *testing.T) {
		key := "teste"
		value := "isso é apenas um teste"
		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		compareError(t, err, errNonExistingWord)
	})

}

func TestUpdate(t *testing.T) {
	key := "teste"
	value := "isso é apenas um teste"
	dictionary := Dictionary{key: value}
	novaDefinicao := "nova definição"

	dictionary.Update(key, novaDefinicao)

	compareValue(t, dictionary, key, novaDefinicao)
}

func TestDelete(t *testing.T) {
	key := "teste"
	dictionary := Dictionary{key: "definição de teste"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	if err != errWordNotFound {
		t.Errorf("espera-se que '%s' seja deletado", key)
	}
}

func compareValue(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()
	result, err := dictionary.Search("teste")

	if err != nil {
		t.Fatal("não foi possível encontrar a palavra adicionada:", err)
	}

	if value != result {
		t.Errorf("resultado '%s', esperado '%s'", result, value)
	}

}

func compareError(t *testing.T, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("resultado erro '%s', esperado '%s'", result, expected)
	}
}

func compareStrings(t *testing.T, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("resultado '%s', expected '%s', dado '%s'", result, expected, "teste")
	}
}
