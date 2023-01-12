package api

import (
	"testing"
)

func TestGetRequester(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := GetRequester("/")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", req)
}
