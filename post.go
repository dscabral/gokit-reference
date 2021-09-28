package post

import (
	"context"
	"errors"
	"time"
)

var (
	// ErrMalformedEntity indicates malformed entity specification (e.g.
	// invalid author or content).
	ErrMalformedEntity = errors.New("malformed entity specification")
	// ErrNotFound indicates a non-existent entity request.
	ErrNotFound = errors.New("non-existent entity")
	// ErrUnsupportedContentType indicates a unsupported content-type (should be application/json)
	ErrUnsupportedContentType = errors.New("unsupported content-type")
)

type Post struct {
	ID       string
	Title    string
	Content  string
	TsCreate time.Time
}

type PostService interface {
	CreatePost(ctx context.Context, post Post) (Post, error)
	ViewPost(ctx context.Context, postID string) (Post, error)
	ViewAllPosts(ctx context.Context) ([]Post, error)
}
