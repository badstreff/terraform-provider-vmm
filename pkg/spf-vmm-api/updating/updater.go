package updating

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
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

func update(doer doer, url string, data []byte) (err error) {
	defer func() {
		if err != nil {
			log.Print(err)
		}
	}()
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("x-ms-principal-id", "terraform-provider-vmm")
	req.Header.Set("Content-Type", "application/json")
	resp, err := doer.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode <= 199 || resp.StatusCode >= 300 {
		return errors.New(string(body))
	}
	return nil
}
