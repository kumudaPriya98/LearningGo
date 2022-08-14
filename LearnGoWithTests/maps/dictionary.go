package main

import "errors"

type Dictionary map[string]string

var ErrNoKey = errors.New("Key not present")

func (d Dictionary) Search(key string) (string, error) {
	value, found := d[key]
	if !found {
		return "", ErrNoKey
	}
	return value, nil
}

func main() {}
