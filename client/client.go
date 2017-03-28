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
	contentTypeXML = "application/xml"
	contentTypeJSON            = "application/json"

	// JSON Rest API
	createCustomCollectionJSON = "/searchblox/rest/v1/api/coladd"
	deleteCustomCollectionJSON = "/searchblox/rest/v1/api/coldelete"
	clearCustomCollectionJSON = "/searchblox/rest/v1/api/clear"

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
	Colname string `json:"colname"`
}

type CustomCollection struct {
	ApiKey   string   `json:"apikey"`
	Document Document `json:"document"`
}

func (s *SearchBlox) CreateCustomCollection(customCollection CustomCollection) error {
	b, err := json.Marshal(customCollection)
	if err != nil {
		return errors.New(encodeErrorJSON)
	}
	url := fmt.Sprintf("%s%s", s.Host, createCustomCollectionJSON)
	resp, err := http.Post(url, contentTypeJSON, bytes.NewBuffer(b))
	if err != nil {
		return errors.New("Create custom collection error")
	}

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

func (s *SearchBlox) DeleteCustomCollection(customCollection CustomCollection) error {
	b, err := json.Marshal(customCollection)
	if err != nil {
		return errors.New(encodeErrorJSON)
	}

	url := fmt.Sprintf("%s%s", s.Host, deleteCustomCollectionJSON)
	resp, err := http.Post(url, contentTypeJSON, bytes.NewBuffer(b))
	if err != nil {
		return errors.New("Delete custom collection error")
	}

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}
