package main

import "errors"

type Dictionary map[string]string

var (
	ErrNoKey        = errors.New("Key not present")
	ErrDuplicateKey = errors.New("Key already present")
)

// the change in the d will be reflected in the callee dictionary instance
// As Map value is "pointer to a runtime.hmap" structure
// When Add is called, the pointer is copied and modified
func (d Dictionary) Add(key string, value string) error {
	_, errSearch := d.Search(key)

	switch errSearch {
	case ErrNoKey:
		d[key] = value
		return nil
	case nil:
		return ErrDuplicateKey
	default:
		return errSearch
	}
}

func (d Dictionary) Search(key string) (string, error) {
	value, found := d[key]
	if !found {
		return "", ErrNoKey
	}
	return value, nil
}

func main() {}
