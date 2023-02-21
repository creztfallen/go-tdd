package maps

type Dictionary map[string]string

type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}

const (
	errWordNotFound    = ErrDictionary("não foi possível encontrar a palavra desejada")
	errExistingKey     = ErrDictionary("não é possível adicionar a palavra pois ela já existe")
	errNonExistingWord = ErrDictionary("não foi possível atualizar a palavra pois ela não existe")
)

func (d Dictionary) Search(key string) (string, error) {
	definition, exists := d[key]
	if !exists {
		return "", errWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case errWordNotFound:
		d[key] = value
	case nil:
		return errExistingKey
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case errWordNotFound:
		return errNonExistingWord
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
