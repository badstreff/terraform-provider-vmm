package deleting

import (
	"errors"
	"fmt"
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

// GenericDeleter generic deleter interface
type GenericDeleter interface {
	DeleteByID(id string) error
}

// NewGenericDeleterService creates a generic deleter service
func NewGenericDeleterService(doer doer, serviceURL string, stamp string) *Service {
	url, err := url.Parse(serviceURL)
	if err != nil {
		return nil
	}
	return &Service{doer: doer, serviceURL: *url, stampID: stamp}
}

// DeleteByID deletes an vmm object by its id
func (s Service) DeleteByID(id string) (err error) {
	uri := fmt.Sprintf(s.serviceURL.String()+"(ID=guid'%s',StampId=guid'%s')", id, s.stampID)
	defer func() {
		if err != nil {
			log.Print(err)
		}
	}()
	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		return err
	}
	req.Header.Set("x-ms-principal-id", "terraform-provider-vmm")
	req.Header.Set("Content-Type", "application/json;odata=minimalmetadata")
	req.Header.Set("Accept", "application/json;odata=minimalmetadata")
	resp, err := s.Do(req)
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
