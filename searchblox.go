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
	// SearchBlox JSON Rest API(https://developer.searchblox.com/v8.6/reference)
	createCustomCollection           = "/searchblox/rest/v1/api/coladd"
	deleteCustomCollection           = "/searchblox/rest/v1/api/coldelete"
	clearCustomCollection            = "/searchblox/rest/v1/api/clear"
	indexDocumentInCustomCollection  = "/searchblox/rest/v1/api/add"
	documentStatusInCustomCollection = "/searchblox/rest/v1/api/status"
	deleteDocumentInCustomCollection = "/searchblox/rest/v1/api/delete"
	addUpdateDocumentInCollection    = "/searchblox/rest/v1/api/docadd"
	deleteDocumentInCollection       = "/searchblox/rest/v1/api/docdelete"
	addCollection                    = "/searchblox/rest/collection/add"
	deleteCollection                 = "/searchblox/rest/collection/delete"
	updatePathInCollection           = "/searchblox/rest/collection/updatePath"
	updateSettingsInCollection       = "/searchblox/rest/collection/updateSettings"
	updateSchedulerInCollection      = "/searchblox/rest/collection/updateScheduler"
	indexStopCrawlerInCollection     = "/searchblox/rest/collection/actions"

	contentType = "application/json"
	encodeError = "JSON encode error"
)

type Client struct {
	Host string
	// FIXME: support protocol and port separately
	// FIXME: store ApiKey in Client struct and use it if no other provided in CustomCollection?
}

// FIXME: maybe need 2 structs: Collection & CustomCollection? And BasicCollection with common fields?
type Collection struct {
	ApiKey            string             `json:"apikey"`
	Document          *Document          `json:"document,omitempty"`
	ColName           string             `json:"colname,omitempty"`
	ColType           string             `json:"coltype,omitempty"`
	RootUrls          []string           `json:"rooturls,omitempty"`
	AllowPaths        []string           `json:"allowpaths,omitempty"`
	DisallowPaths     []string           `json:"disallowpaths,omitempty"`
	AllowFormat       []string           `json:"allowformat,omitempty"`
	KeywordInContext  string             `json:"keyword-in-context,omitempty"`
	RemoveDuplicates  string             `json:"remove-duplicates,omitempty"`
	Boost             string             `json:"boost,omitempty"`
	Stemming          string             `json:"stemming,omitempty"`
	Spelling          string             `json:"spelling,omitempty"`
	Logging           string             `json:"logging,omitempty"`
	HtmlSettings      *HtmlSettings      `json:"html-settings,omitempty"`
	BasicAuthSettings *BasicAuthSettings `json:"basic-auth-settings,omitempty"`
	FormAuthSettings  *FormAuthSettings  `json:"form-auth-settings,omitempty"`
	ProxySettings     *ProxySettings     `json:"proxy-settings,omitempty"`
	Index             *Scheduler         `json:"index,omitempty"`
	Clear             *Scheduler         `json:"clear,omitempty"`
	Refresh           *Scheduler         `json:"refresh,omitempty""`
	Action            string             `json:"action,omitempty"`
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

type HtmlSettings struct {
	Description    string `json:"description,omitempty"`
	MaxDocAge      string `json:"max-doc-age,omitempty"`
	MaxDocSize     string `json:"max-doc-size,omitempty"`
	SpiderMaxDepth string `json:"spider-max-depth,omitempty"`
	SpiderMaxDelay string `json:"spider-max-delay,omitempty"`
	UserAgent      string `json:"user-agent,omitempty"`
	Referer        string `json:"referer,omitempty"`
	IgnoreRobots   string `json:"ignore-robots,omitempty"`
	FollowSitemap  string `json:"follow-sitemap,omitempty"`
	FollowRedirect string `json:"follow-redirect,omitempty"`
}

type BasicAuthSettings struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type FormAuthSettings struct {
	FormUrl    string `json:"form-url,omitempty"`
	FormAction string `json:"form-action,omitempty"`
	Form       []Form `json:"form,omitempty"`
}

type Form struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ProxySettings struct {
	ServerUrl string `json:"server-url,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

type Scheduler struct {
	Frequency string `json:"frequency,omitemtpy"` // FIXME: Do we need to omit these values for all nested structs?
	TimeStamp string `json:"timestamp,omitemtpy"`
}

//FIXME: better searchblox exception handler(Bad request, etc)
/*
General method to make a call to SearchBlox Rest API with Collection struct
*/
func (s *Client) makeCall(url string, collection Collection) ([]byte, error) {
	b, err := json.Marshal(collection)
	if err != nil {
		return nil, errors.New(encodeError)
	}
	fmt.Println("\n\n\nrequest Url: ", url)
	fmt.Print(string(b))
	resp, err := http.Post(url, contentType, bytes.NewBuffer(b))
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

/*
CREATE CUSTOM COLLECTION
https://developer.searchblox.com/v8.6/reference#restv1apicoladd
*/
func (s *Client) CreateCustomCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, createCustomCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
DELETE CUSTOM COLLECTION
https://developer.searchblox.com/v8.6/reference#restv1apicoldelete
*/
func (s *Client) DeleteCustomCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, deleteCustomCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
CLEAR CUSTOM COLLECTION
https://developer.searchblox.com/v8.6/reference#json-restv1apiclear
*/
func (s *Client) ClearCustomCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, clearCustomCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
INDEX DOCUMENT - CUSTOM COLLECTION
https://developer.searchblox.com/v8.6/reference#restv1apicoldelete-1
*/
func (s *Client) IndexDocumentInCustomCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, indexDocumentInCustomCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
DOCUMENT STATUS - CUSTOM COLLECTION
https://developer.searchblox.com/v8.6/reference#json-restv1apistatus
*/
func (s *Client) DocumentStatusInCustomCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, documentStatusInCustomCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
DELETE DOCUMENT - CUSTOM COLLECTION
https://developer.searchblox.com/v8.6/reference#json-restv1apidelete
*/
func (s *Client) DeleteDocumentInCustomCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, deleteDocumentInCustomCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
ADD/UPDATE DOCUMENT - HTTP/FILESYSTEM COLLECTION
https://developer.searchblox.com/v8.6/reference#json-restv1apidocadd
*/
func (s *Client) AddUpdateDocumentInCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, addUpdateDocumentInCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
DELETE DOCUMENT - HTTP/FILESYSTEM COLLECTION
https://developer.searchblox.com/v8.6/reference#json-restv1apidocdelete
*/
func (s *Client) DeleteDocumentInCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, deleteDocumentInCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
ADD HTTP, FILESYSTEM OR DATABASE COLLECTION
https://developer.searchblox.com/v8.6/reference#json-restcollectionadd
*/
func (s *Client) AddCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, addCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
DELETE HTTP, FILESYSTEM OR DATABASE COLLECTION
https://developer.searchblox.com/v8.6/reference#json-restcollectiondelete
*/
func (s *Client) DeleteCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, deleteCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
UPDATE PATH - HTTP, FILESYSTEM OR DATABASE COLLECTION
https://developer.searchblox.com/v8.6/reference#json-restcollectionupdate
*/
func (s *Client) UpdatePathInCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, updatePathInCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
UPDATE SETTINGS - HTTP, FILESYSTEM OR DATABASE COLLECTION
https://developer.searchblox.com/v8.6/reference#json-restcollectionupdatesettings
*/
func (s *Client) UpdateSettingsInCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, updateSettingsInCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
UPDATE SCHEDULER - HTTP, FILESYSTEM OR DATABASE COLLECTION
https://developer.searchblox.com/v8.6/reference#json-restcollectionupdatescheduler
*/
func (s *Client) UpdateSchedulerInCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, updateSchedulerInCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
INDEX/STOP CRAWLER - HTTP, FILESYSTEM OR DATABASE COLLECTION
https://developer.searchblox.com/v8.6/reference#json-restcollectionactions
*/
func (s *Client) IndexStopCrawlerInCollection(collection Collection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, indexStopCrawlerInCollection)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
