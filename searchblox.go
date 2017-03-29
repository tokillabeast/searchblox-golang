package searchblox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	contentTypeJSON = "application/json"

	// SearchBlox JSON Rest API(https://developer.searchblox.com/v8.6/reference)
	createCustomCollectionJSON                             = "/searchblox/rest/v1/api/coladd"
	deleteCustomCollectionJSON                             = "/searchblox/rest/v1/api/coldelete"
	clearCustomCollectionJSON                              = "/searchblox/rest/v1/api/clear"
	indexDocumentCustomCollectionJSON                      = "/searchblox/rest/v1/api/add"
	documentStatusCustomCollectionJSON                     = "/searchblox/rest/v1/api/status"
	deleteDocumentCustomCollectionJSON                     = "/searchblox/rest/v1/api/delete"
	addUpdateDocumentHttpFileSystemCollectionJSON          = "/searchblox/rest/v1/api/docadd"
	deleteDocumentHttpFileSystemCollectionJSON             = "/searchblox/rest/v1/api/docdelete"
	addHttpFileSystemOrDatabaseCollectionJSON              = "/searchblox/rest/collection/add"
	deleteHttpFileSystemOrDatabaseCollectionJSON           = "/searchblox/rest/collection/delete"
	updatePathHttpFileSystemOrDatabaseCollectionJSON       = "/searchblox/rest/collection/updatePath"
	updateSettingsHttpFileSystemOrDatabaseCollectionJSON   = "/searchblox/rest/collection/updateSettings"
	updateSchedulerHttpFileSystemOrDatabaseCollectionJSON  = "/searchblox/rest/collection/updateScheduler"
	indexStopCrawlerHttpFileSystemOrDatabaseCollectionJSON = "/searchblox/rest/collection/actions"

	// errors
	encodeErrorJSON = "JSON encode error"
)

type Client struct {
	Host string
	// FIXME: support protocol and port separately
	// FIXME: support XML format
	// FIXME: store ApiKey in Client struct?
}

type Document struct {
	ColName      string            `json:"colname"`
	Url          string            `json:"url,omitempty"`
	Uid          string            `json:"uid,omitempty"`
	Location     string            `json:"location,omitempty"`
	Alpha        string            `json:"alpha,omitempty"`
	Size         int               `json:"size,omitempty"`
	Title        string            `json:"title,omitempty"`
	Keywords     string            `json:"keywords,omitempty"`
	Description  string            `json:"description,omitempty"`
	Content      string            `json:"content,omitempty"`
	LastModified string            `json:"lastmodified,omitempty"`
	ContentType  string            `json:"contenttype,omitempty"`
	Meta         map[string]string `json:"meta"`
}

type CustomCollection struct {
	ApiKey   string   `json:"apikey"`
	Document Document `json:"document"`
}

//FIXME: better searchblox exception handler(Bad request, etc)
func (s *Client) makeCall(url string, customCollection CustomCollection) ([]byte, error) {
	b, err := json.Marshal(customCollection)
	if err != nil {
		return nil, errors.New(encodeErrorJSON)
	}
	fmt.Println("\n\n\nrequest Url: ", url)
	fmt.Print(string(b))
	resp, err := http.Post(url, contentTypeJSON, bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.New("Custom collection error")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status: %s\n Body: %s", resp.Status, resp.Body)
	}
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response StatusCode:", resp.StatusCode)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return body, nil
}

func (s *Client) CreateCustomCollection(customCollection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, createCustomCollectionJSON)
	body, err := s.makeCall(url, customCollection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) DeleteCustomCollection(customCollection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, deleteCustomCollectionJSON)
	body, err := s.makeCall(url, customCollection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) ClearCustomCollection(customCollection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, clearCustomCollectionJSON)
	body, err := s.makeCall(url, customCollection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) IndexDocumentCustomCollection(customCollection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, indexDocumentCustomCollectionJSON)
	body, err := s.makeCall(url, customCollection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) DocumentStatusCustomCollection(customCollection CustomCollection) (string, error) { //FIXME: return map to support meta?
	url := fmt.Sprintf("%s%s", s.Host, documentStatusCustomCollectionJSON)
	body, err := s.makeCall(url, customCollection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
