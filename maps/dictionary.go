package maps

// Dictionary store definitions to word
type Dictionary map[string]string

// DictionaryErr are errors that can happen when interaction with the dictionary
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	// ErrNotFound means the definition could not be found for the given word
	ErrNotFound = DictionaryErr("not found")

	// ErrWordExists means you are trying to add a word that is already known
	ErrWordExists = DictionaryErr("word already exists")

	// ErrWordDoesNotExist means you are trying to update a word that does not exist
	ErrWordDoesNotExist = DictionaryErr("word doesn't exist")
)

// Add inserts a word and definition into the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Delete removes a word from a dictionary
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}

// Search find a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Update changes the definition of a given word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}
