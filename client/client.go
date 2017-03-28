package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	contentTypeXML = "application/xml"

	contentTypeJSON            = "application/json"
	createCustomCollectionJSON = "/searchblox/rest/v1/api/coladd"
	deleteCustomCollectionJSON = "/searchblox/rest/v1/api/coldelete"
)

type SearchBlox struct {
	Host string
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
		return errors.New("JSON encode error")
	}
	url := fmt.Sprintf("%s%s", s.Host, createCustomCollectionJSON)
	_, err = http.Post(url, contentTypeJSON, bytes.NewBuffer(b))
	if err != nil {
		return errors.New("Create custom collection error")
	}
	return nil
}
