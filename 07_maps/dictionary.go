package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(word string) (string, error) {

	value, ok := dictionary[word]

	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (dictionary Dictionary) Add(word string, definition string) error {

	_, err := dictionary.Search(word)

	switch err {
		case ErrNotFound:
			dictionary[word] = definition

		case nil:
			return ErrWordExists

		default:
			return err
	}

	return nil
}


func (dictionary Dictionary) Update(word, definition string) error {

	_, err := dictionary.Search(word)

	switch err {
		case ErrNotFound:
			return ErrWordDoesNotExist

		case nil:
			dictionary[word] = definition
			return nil

		default:
			return err
	}
}

func (dictionary Dictionary) Delete(word string) {
	delete(dictionary, word)
}