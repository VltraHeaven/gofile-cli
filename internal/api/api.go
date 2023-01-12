package api

import (
	"io"
	"net/http"
)

const api string = "https://api.gofile.io"

func GetRequester(req *Request) (b []byte, err error) {
	path := api + req.Endpoint
	if req.Token != "" {
		path += "?token=" + req.Token
	}

	if req.ExtraParams != nil {
		for k, v := range req.ExtraParams {
			path += "&" + k + "=" + v
		}
	}

	res, err := http.Get(path)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	b, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}

func NewGetRequest(endpoint string, token string, params map[string]string) *Request {
	return &Request{
		Endpoint:    endpoint,
		Token:       token,
		ExtraParams: params,
	}
}

func PutRequester(req *Request) (b []byte, err error) {
	return
}
