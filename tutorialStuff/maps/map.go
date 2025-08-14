package maps

const (
	ErrNotInDictionary   = DictionaryErr("Needle doesn't exist")
	ErrAlreadyInDictionary = DictionaryErr("Can not add, already exists in dictionary")
	ErrUpdatedKeyDoesntExist = DictionaryErr("Can not update key, it does not exist in dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(needle string) (string, DictionaryErr) {
	value, err := d[needle]

	if !err {
		return "", ErrNotInDictionary
	}
	return value, ""
}

func (d Dictionary) Add(key, value string) (DictionaryErr) {
	_, err := d.Search(key)

	switch err {
		case ErrNotInDictionary:
			d[key] = value
		case "":
			return ErrAlreadyInDictionary
		default:
			return err
	}

	return ""
}

func (d Dictionary) Update(key, value string) (DictionaryErr) {
	_, err := d.Search(key)

	switch err {
		case ErrNotInDictionary:
			return ErrUpdatedKeyDoesntExist
		case "":
			d[key] = value
		default:
			return err
	}

	return ""
}