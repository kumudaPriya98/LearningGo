package main

import "errors"

type Dictionary map[string]string

var ErrNoKey = errors.New("Key not present")

// the change in the d will be reflected in the callee dictionary instance
// As Map value is "pointer to a runtime.hmap" structure
// When Add is called, the pointer is copied and modified
func (d Dictionary) Add(key string, value string) {
	d[key] = value
}

func (d Dictionary) Search(key string) (string, error) {
	value, found := d[key]
	if !found {
		return "", ErrNoKey
	}
	return value, nil
}

func main() {}
