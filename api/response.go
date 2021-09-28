package api

import "net/http"

type postRes struct {
	ID       string `json:"id"`
	Title    string `json:"author"`
	Content  string `json:"content"`
	TsCreate string `json:"ts_created"`
	created  bool
}

func (s postRes) Code() int {
	if s.created {
		return http.StatusCreated
	}

	return http.StatusOK
}

func (s postRes) Headers() map[string]string {
	return map[string]string{}
}

func (s postRes) Empty() bool {
	return false
}
