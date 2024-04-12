package loader

import "errors"

var (
	// ErrLoaderInternal is an error that returns when a internal error occurs
	ErrLoaderInternal = errors.New("loader: internal error")
)

// Loader is an interface that contains a method that loads the entities into bd
type Loader interface {
	// Load is a method that loads the entities into bd
	Load() (err error)
}
