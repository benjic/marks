package lib

import "net/url"

// An Entry represents a resource.
type Entry interface {
	// ID provides a way to uniformly identify an Entry.
	ID() string

	// URL provides the resource location of an Entry.
	URL() *url.URL
}

type entry struct {
	u *url.URL
}

func (e *entry) ID() string    { return e.u.String() }
func (e *entry) URL() *url.URL { return e.u }

// NewEntry creates a new Entry for the given URI string.
func NewEntry(value string) (Entry, error) {
	u, err := url.Parse(value)

	if err != nil {
		return nil, err
	}

	return &entry{u}, nil
}
