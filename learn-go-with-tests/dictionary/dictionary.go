package dictionary

import "errors"

var KeyNotFoundError = errors.New("key not found")
var DuplicateKeyError = errors.New("key already exists")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
    if val, ok := d[key]; ok {
        return val, nil
    } else {
        return "", KeyNotFoundError
    }
}

func (d Dictionary) Add(key string, value string) error {
    if _, ok := d[key]; ok {
        return DuplicateKeyError
    }
    d[key] = value
    return nil
}

func (d Dictionary) Update(key string, value string) error {
    if _, ok := d[key]; !ok {
        return KeyNotFoundError
    }
    d[key] = value
    return nil
}
