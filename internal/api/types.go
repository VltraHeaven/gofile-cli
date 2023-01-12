package api

const api string = "https://api.gofile.io"

type Request struct {
	Endpoint, Token string
	ExtraParams     map[string]string
}
