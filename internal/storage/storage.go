package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(user string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

var ErrNoSavedPages = errors.New("no saved pages")

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()
	if _, err := h.Write([]byte(p.URL)); err != nil {
		return "", fmt.Errorf("hash error: %w", err)
	}
	if _, err := h.Write([]byte(p.UserName)); err != nil {
		return "", fmt.Errorf("hash error: %w", err)
	}

	return string(h.Sum(nil)), nil
}
