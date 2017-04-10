package lib

import (
	"bufio"
	"io"
	"strings"
)

// A Provider is a source of URLs.
type Provider interface {
	// Add includes a given value.
	Add(value Entry) error

	// Suggest takes an input value and returns a list of suggest values.
	Suggest(terms []string) []Entry
}

type provider struct {
	io.ReadWriter
	entries map[string]Entry
}

func (p *provider) Add(value Entry) error {
	if _, ok := p.entries[value.ID()]; !ok {
		p.entries[value.ID()] = value
		_, err := p.Write([]byte(value.URL().String() + "\r\n"))

		return err
	}

	return nil
}

func (p *provider) Suggest(terms []string) []Entry {
	results := make([]Entry, 0)

	for _, e := range p.entries {
		u := e.URL().String()

		if matchesAllTerms(u, terms) {
			results = append(results, e)
		}
	}

	return results
}

// NewProvider establishes a provider which draws entries from the given
// filepath.
func NewProvider(rw io.ReadWriter) (Provider, error) {
	scanner := bufio.NewScanner(rw)
	entries := make(map[string]Entry, 0)

	for scanner.Scan() {
		entry, err := NewEntry(scanner.Text())

		if err != nil {
			return nil, err
		}

		entries[entry.ID()] = entry
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &provider{rw, entries}, nil
}

func matchesAllTerms(u string, terms []string) bool {
	for _, term := range terms {
		if !strings.Contains(u, term) {
			return false
		}
	}
	return true
}
