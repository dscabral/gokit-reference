package api

import (
	"context"
	"time"

	post "github.com/dcabral/gokit"
	"github.com/go-kit/kit/endpoint"
)

func addPostEnpoint(svc post.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		post := post.Post{
			ID:       req.ID,
			Title:    req.Title,
			Content:  req.Content,
			TsCreate: time.Now(),
		}

		saved, err := svc.CreatePost(ctx, post)
		if err != nil {
			return nil, err
		}

		res := postRes{
			ID:       saved.ID,
			Title:    saved.Title,
			Content:  saved.Content,
			TsCreate: saved.TsCreate.String(),
			created:  true,
		}

		return res, nil

	}
}

func viewPostEnpoint(svc post.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(viewResourceReq)
		if err := req.validate(); err != nil {
			return nil, err
		}

		post, err := svc.ViewPost(ctx, req.id)
		if err != nil {
			return nil, err
		}

		res := postRes{
			ID:       post.ID,
			Title:    post.Title,
			Content:  post.Content,
			TsCreate: post.TsCreate.String(),
		}
		return res, nil
	}
}

func viewAllPostsEnpoint(svc post.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		list, err := svc.ViewAllPosts(ctx)
		if err != nil {
			return nil, err
		}
		var res = []postRes{}
		for _, v := range list {
			res = append(res, postRes{
				ID:       v.ID,
				Title:    v.Title,
				Content:  v.Content,
				TsCreate: v.TsCreate.String(),
			})
		}
		return res, nil
	}
}
