package api

import (
	"testing"
)

func TestGetRequester(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req := NewGetRequest("/", "", nil)
	res, err := GetRequester(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("%s", res)
}

func TestPutRequester(t *testing.T) {
	req, err := PutRequester(nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("%s", req)
}
