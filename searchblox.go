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

type CustomCollection struct {
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
func (s *Client) makeCall(url string, collection CustomCollection) ([]byte, error) {
	b, err := json.Marshal(collection)
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

func (s *Client) CreateCustomCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, createCustomCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) DeleteCustomCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, deleteCustomCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) ClearCustomCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, clearCustomCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) IndexDocumentCustomCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, indexDocumentCustomCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) DocumentStatusCustomCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, documentStatusCustomCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) DeleteDocumentCustomCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, deleteDocumentCustomCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) AddUpdateDocumentHttpFileSystemCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, addUpdateDocumentHttpFileSystemCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) DeleteDocumentHttpFileSystemCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, deleteDocumentHttpFileSystemCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) AddHttpFileSystemOrDatabaseCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, addHttpFileSystemOrDatabaseCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) DeleteHttpFileSystemOrDatabaseCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, deleteHttpFileSystemOrDatabaseCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) UpdatePathHttpFileSystemOrDatabaseCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, updatePathHttpFileSystemOrDatabaseCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) UpdateSettingsHttpFileSystemOrDatabaseCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, updateSettingsHttpFileSystemOrDatabaseCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (s *Client) UpdateSchedulerHttpFileSystemOrDatabaseCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, updateSchedulerHttpFileSystemOrDatabaseCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// FIXME: more simple titles? Maybe IndexStopCrawler and comments for this method
// INDEX/STOP CRAWLER - HTTP, FILESYSTEM OR DATABASE COLLECTION
func (s *Client) IndexStopCrawlerHttpFileSystemOrDatabaseCollection(collection CustomCollection) (string, error) {
	url := fmt.Sprintf("%s%s", s.Host, indexStopCrawlerHttpFileSystemOrDatabaseCollectionJSON)
	body, err := s.makeCall(url, collection)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
