package dictionary

const (
	ErrNotFound         = DictionaryErr("could not find word")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (dic Dictionary) Search(text string) (string, error) {
	definition, ok := dic[text]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (dic Dictionary) Add(word, definition string) error {
	_, err := dic.Search(word)

	switch err {
	case ErrNotFound:
		dic[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dic Dictionary) Update(word string, definition string) error {
	_, err := dic.Search(word)

	if err != nil {
		return ErrWordDoesNotExist
	}

	dic[word] = definition
	return nil
}

func (dic Dictionary) Delete(word string) {
	delete(dic, word)
}
