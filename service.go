package post

import (
	"go.uber.org/zap"
	"sync"
)

var _ PostService = (*postService)(nil)

type postService struct {
	logger *zap.Logger
	post   map[string]Post
	mu     sync.Mutex
}

func NewService(logger *zap.Logger) PostService {
	return &postService{
		logger: logger,
		post:   make(map[string]Post),
	}
}
