package post

import (
	"context"
	"github.com/gofrs/uuid"
)

func (svc *postService) CreatePost(ctx context.Context, post Post) (Post, error) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	ID, err := uuid.NewV4()
	if err != nil {
		return Post{}, err
	}
	post.ID = ID.String()
	svc.post[post.ID] = post
	return post, nil
}

func (svc *postService) ViewPost(ctx context.Context, postID string) (Post, error) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	if c, ok := svc.post[postID]; ok {
		return c, nil
	}

	return Post{}, ErrNotFound
}

func (svc *postService) ViewAllPosts(ctx context.Context) ([]Post, error) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	var posts = []Post{}
	for _, v := range svc.post {
		posts = append(posts, v)
	}
	return posts, nil
}
