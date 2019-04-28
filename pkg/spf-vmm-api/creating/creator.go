package creating

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// Service getter service
type Service struct {
	Doer
	serviceURL url.URL
	stampID    string
}

func create(doer Doer, url string, data []byte) (id *string, err error) {
	defer func() {
		if err != nil {
			log.Print(err)
		}
	}()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-ms-principal-id", "terraform-provider-vmm")
	req.Header.Set("Content-Type", "application/json")
	resp, err := doer.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode <= 199 || resp.StatusCode >= 300 {
		return nil, errors.New(string(body))
	}
	result := struct {
		ID *string `xml:"content>properties>ID" json:"-"`
	}{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result.ID, nil
}
