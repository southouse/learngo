package mydict

import "errors"

// Dictionary Type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Can`t update non-existing word")
	errCantDelete = errors.New("Can`t delete non-existing word")
	errExists     = errors.New("That word is already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}

	return "", errNotFound
}

// Add for a word
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errExists
	}

	return nil
}

// Update for a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		d[word] = definition
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errCantDelete
	case nil:
		delete(d, word)
	}

	return nil
}
