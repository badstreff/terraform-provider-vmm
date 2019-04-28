package getting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type doer interface {
	Do(*http.Request) (*http.Response, error)
}

// Service getter service
type Service struct {
	doer
	serviceURL url.URL
	stampID    string
}

func get(doer doer, uri string) ([]byte, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := doer.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return nil, fmt.Errorf("Error: Non 2xx Status Code\n Response Body: %s", string(body))
	}
	return body, nil
}
