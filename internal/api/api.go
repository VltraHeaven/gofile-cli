package api

import (
	"io"
	"net/http"
)

const api string = "https://api.gofile.io"

func GetRequester(endpoint string) (b []byte, err error) {
	path := api + endpoint
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
