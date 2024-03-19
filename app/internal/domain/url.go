package domain

import "errors"

// maybe we can count on url client ?
type MyUrl struct {
	Slug string
	Url  string
}

var ErrNoResult = errors.New("no result")
