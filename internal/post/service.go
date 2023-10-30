package post

import (
	"errors"
	"unicode/utf8"

	"github.com/ogabriel/ytgoapi/internal"

	"github.com/google/uuid"
)

var ErrPostBodyEmpty = errors.New("post body is empty")
var ErrPostBodyExceedsLimit = errors.New("post body exceeds limit")
var ErrPostNotFound = errors.New("post not found")

type Service struct {
	Repository Repository
}

func (p Service) Create(post internal.Post) error {
	if post.Body == "" {
		return ErrPostBodyEmpty
	}

	if utf8.RuneCountInString(post.Body) > 140 {
		return ErrPostBodyExceedsLimit
	}

	return p.Repository.Insert(post)
}

func (p Service) FindOneById(id uuid.UUID) (internal.Post, error) {
	return p.Repository.FindOneById(id)
}

func (p Service) Delete(id uuid.UUID) error {
	return p.Repository.Delete(id)
}
