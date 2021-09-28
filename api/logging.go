package api

import (
	"context"
	post "github.com/dcabral/gokit"
	"go.uber.org/zap"
	"time"
)

var _ post.PostService = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger *zap.Logger
	svc    post.PostService
}

func (l loggingMiddleware) CreatePost(ctx context.Context, post post.Post) (_ post.Post, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Warn("method call: create_post",
				zap.Error(err),
				zap.Duration("duration", time.Since(begin)))
		} else {
			l.logger.Info("method call: create_post",
				zap.Duration("duration", time.Since(begin)))
		}
	}(time.Now())
	return l.svc.CreatePost(ctx, post)
}

func (l loggingMiddleware) ViewPost(ctx context.Context, postID string) (_ post.Post, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Warn("method call: view_post",
				zap.Error(err),
				zap.Duration("duration", time.Since(begin)))
		} else {
			l.logger.Info("method call: view_post",
				zap.Duration("duration", time.Since(begin)))
		}
	}(time.Now())
	return l.svc.ViewPost(ctx, postID)
}

func (l loggingMiddleware) ViewAllPosts(ctx context.Context) (_ []post.Post, err error) {
	defer func(begin time.Time) {
		if err != nil {
			l.logger.Warn("method call: view_all_posts",
				zap.Error(err),
				zap.Duration("duration", time.Since(begin)))
		} else {
			l.logger.Info("method call: view_all_posts",
				zap.Duration("duration", time.Since(begin)))
		}
	}(time.Now())
	return l.svc.ViewAllPosts(ctx)
}

func NewLoggingMiddleware(svc post.PostService, logger *zap.Logger) post.PostService {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
