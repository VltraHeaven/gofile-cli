package api

type Request struct {
	Endpoint, Token string
	ExtraParams     map[string]string
}
