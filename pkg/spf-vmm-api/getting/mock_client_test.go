package getting

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type MockDoer struct {
	Data       string
	StatusCode int
}

func (d MockDoer) Do(*http.Request) (*http.Response, error) {
	resp := http.Response{
		Body:       ioutil.NopCloser(bytes.NewBufferString(d.Data)),
		StatusCode: d.StatusCode,
	}
	return &resp, nil
}
