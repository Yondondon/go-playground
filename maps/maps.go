package maps

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrNotFound    = DictionaryErr("could not find the key you were looking for")
	ErrKeyExists   = DictionaryErr("such key already exists")
	ErrKeyNotFound = DictionaryErr("couldn't perform operation on key because it doesn't exist n ''")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	result, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrKeyNotFound
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrKeyNotFound
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}
