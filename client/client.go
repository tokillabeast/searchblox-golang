package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	contentTypeXML  = "application/xml"
	contentTypeJSON = "application/json"

	// JSON Rest API
	createCustomCollectionJSON         = "/searchblox/rest/v1/api/coladd"
	deleteCustomCollectionJSON         = "/searchblox/rest/v1/api/coldelete"
	clearCustomCollectionJSON          = "/searchblox/rest/v1/api/clear"
	indexDocumentCustomCollectionJSON  = "/searchblox/rest/v1/api/add"
	documentStatusCustomCollectionJSON = "/searchblox/rest/v1/api/status"
	deleteDocumentCustomCollectionJSON = "/searchblox/rest/v1/api/delete"

	// errors
	encodeErrorJSON = "JSON encode error"
)

type SearchBlox struct {
	Host string
	// FIXME: support protocol and port separately
	// FIXME: support XML format
	// FIXME: store ApiKey in Client struct?
}

type Document struct {
	Colname      string `json:"colname"`
	Url          string `json:"url,omitempty"`
	Uid          string `json:"uid,omitempty"`
	Location     string `json:"location,omitempty"`
	Alpha        string `json:"alpha,omitempty"`
	Size         string `json:"size,omitempty"`
	Title        string `json:"title,omitempty"`
	Keywords     string `json:"keywords,omitempty"`
	Description  string `json:"description,omitempty"`
	Content      string `json:"content,omitempty"`
	LastModified string `json:"lastmodified,omitempty"`
	ContentType  string `json:"contenttype,omitempty"`
	meta         struct {
		Location string `json:"location.omitempty"`
		Temp     string `json:"temp,omitempty"`
		Weather  string `json:"weather,omitempty"`
	}
}

type CustomCollection struct {
	ApiKey   string   `json:"apikey"`
	Document Document `json:"document"`
}

func (s *SearchBlox) makeCustomCollectionCall(url string, customCollection CustomCollection) error {
	b, err := json.Marshal(customCollection)
	if err != nil {
		return errors.New(encodeErrorJSON)
	}
	resp, err := http.Post(url, contentTypeJSON, bytes.NewBuffer(b))
	if err != nil {
		return errors.New("Custom collection error")
	}

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

func (s *SearchBlox) CreateCustomCollection(customCollection CustomCollection) error {
	url := fmt.Sprintf("%s%s", s.Host, createCustomCollectionJSON)
	err := s.makeCustomCollectionCall(url, customCollection)
	if err != nil {
		return err
	}
	return nil
}

func (s *SearchBlox) DeleteCustomCollection(customCollection CustomCollection) error {
	url := fmt.Sprintf("%s%s", s.Host, deleteCustomCollectionJSON)
	err := s.makeCustomCollectionCall(url, customCollection)
	if err != nil {
		return err
	}
	return nil
}

func (s *SearchBlox) ClearCustomCollection(customCollection CustomCollection) error {
	url := fmt.Sprintf("%s%s", s.Host, clearCustomCollectionJSON)
	err := s.makeCustomCollectionCall(url, customCollection)
	if err != nil {
		return err
	}
	return nil
}
