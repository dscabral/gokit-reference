package api

import (
	post "github.com/dcabral/gokit"
	"github.com/mainflux/mainflux/pkg/errors"
)

type addReq struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content,omitempty"`
}

func (req addReq) validate() error {
	if req.Title == "" || req.Content == "" {
		return post.ErrMalformedEntity
	}
	return nil
}

type viewResourceReq struct {
	token string
	id    string
}

func (req viewResourceReq) validate() error {
	if req.id == "" {
		return errors.ErrMalformedEntity
	}
	return nil
}

type viewListResourceReq struct {
}
