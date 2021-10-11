package api

import (
	"context"
	"encoding/json"
	post "github.com/dcabral/gokit"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-zoo/bone"
	"net/http"
	"strings"
)

func MakeHandler(svcName string, svc post.PostService) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	r := bone.New()

	r.Post("/posts", kithttp.NewServer(
		addPostEnpoint(svc),
		decodeAddRequest,
		EncodeResponse,
		opts...,
	))
	r.Get("/posts/:id", kithttp.NewServer(
		viewPostEnpoint(svc),
		decodeView,
		EncodeResponse,
		opts...,
	))
	r.Get("/posts", kithttp.NewServer(
		viewAllPostsEnpoint(svc),
		decodeViewList,
		EncodeResponse,
		opts...,
	))
	r.GetFunc("/live", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("UP"))
	})

	return r

}

func decodeViewList(_ context.Context, r *http.Request) (interface{}, error) {
	req := viewListResourceReq{}
	return req, nil
}

func decodeView(_ context.Context, r *http.Request) (interface{}, error) {
	req := viewResourceReq{
		id: bone.GetValue(r, "id"),
	}
	return req, nil
}

func decodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		return nil, post.ErrUnsupportedContentType
	}
	req := addReq{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, post.ErrMalformedEntity
	}
	return req, nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	switch err {
	case post.ErrNotFound:
		w.WriteHeader(http.StatusNotFound)
	case post.ErrMalformedEntity:
		w.WriteHeader(http.StatusBadRequest)
	case post.ErrUnsupportedContentType:
		w.WriteHeader(http.StatusUnsupportedMediaType)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
