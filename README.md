# golang-todo

### Overview

Simple Golang Client for SearchBlox API(v8.6) - https://www.searchblox.com/
Used technology stack for development:
- Go(1.8.0) programming language - https://golang.org/
- Git(2.9.3) - https://git-scm.com/
- Docker(17.03.0-ce) - https://github.com/docker/docker
- Docker-Compose(1.11.2) - https://github.com/docker/compose

To run SearchBlox container Docker and Docker-Compose should be installed locally.
<br /> Docker install: https://docs.docker.com/engine/installation/
<br /> Docker-Compose install: https://docs.docker.com/compose/install/

### Usage
To use this package open terminal and execute(or use some golang package manager):
```
    go get github.com/tokillamockingbird/searchblox-golang
```
To start SearchBox locally with Docker just locate to `searchblox-golang` directory and execute command:
```
    docker-compose start searchblox
```
After this you can check Searchblox on 80 port, open `localhost:80` in browser to check it.

```go
package main

import (
	"github.com/tokillamockingbird/searchblox-golang"
)

func main() {
	apiKey := "47AD645E72CEA9D8B2AB08A2312BF432"
	c := searchblox.Client{Host: "http://localhost:80"}

	customCollection := searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName: "Test",
		},
	}
	// FIXME: check if we have limit on collection creation or error during creation situations
	_, err := c.CreateCustomCollection(customCollection)
	if err != nil {
		panic(err)
	}

	indexDocument := searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName:      "Test",
			Url:          "http://www.searchblox.com",
			Uid:          "http://www.searchblox.com",
			Location:     "http://www.searchblox.com",
			Alpha:        "string",
			Size:         44244,
			Title:        "Text",
			Keywords:     "keywords",
			Description:  "SearchBlox Content Search Software",
			Content:      "content",
			LastModified: "14 January 2015 06:19:42 GMT",
			ContentType:  "HTML",
			Meta: map[string]string{
				"custom": "Value",
			},
		},
	}
	_, err = c.IndexDocumentInCustomCollection(indexDocument)
	if err != nil {
		panic(err)
	}

	documentStatus := searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName: "Test",
			Uid:     "http://www.searchblox.com",
		},
	}
	_, err = c.DocumentStatusInCustomCollection(documentStatus)
	if err != nil {
		panic(err)
	}

	deleteDocument := searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName: "Test",
			Uid:     "http://www.searchblox.com", // Maybe Location?
		},
	}
	_, err = c.DeleteDocumentInCustomCollection(deleteDocument)
	if err != nil {
		panic(err)
	}

	_, err = c.ClearCustomCollection(customCollection)
	if err != nil {
		panic(err)
	}

	_, err = c.DeleteCustomCollection(customCollection)
	if err != nil {
		panic(err)
	}

	addUpdateDocument := searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName:  "Test",
			Location: "http://www.searchblox.com",
		},
	}
	_, err = c.AddUpdateDocumentInCollection(addUpdateDocument)
	if err != nil {
		panic(err)
	}

	deleteDocument = searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName:  "Test",
			Location: "http://www.searchblox.com",
		},
	}
	_, err = c.DeleteDocumentInCollection(deleteDocument)
	if err != nil {
		panic(err)
	}

	addCollection := searchblox.Collection{
		ApiKey:  apiKey,
		ColName: "httpcollection",
		ColType: "http",
	}
	_, err = c.AddCollection(addCollection)
	if err != nil {
		panic(err)
	}

	deleteCollection := searchblox.Collection{
		ApiKey:  apiKey,
		ColName: "httpcollection",
	}
	_, err = c.DeleteCollection(deleteCollection)
	if err != nil {
		panic(err)
	}

	updatePath := searchblox.Collection{
		ApiKey:  apiKey,
		ColName: "httpcollection",
		RootUrls: []string{
			"http://www.google.co.in",
			"http://www.bing.com",
		},
		AllowPaths: []string{
			".*",
		},
		DisallowPaths: []string{
			"http://www.google.co.in/test/bingo",
		},
		AllowFormat: []string{
			"HTML",
			"text",
		},
	}
	_, err = c.UpdatePathInCollection(updatePath)
	if err != nil {
		panic(err)
	}

	updateSettings := searchblox.Collection{
		ApiKey:           apiKey,
		ColName:          "httpcollection",
		KeywordInContext: "false",
		RemoveDuplicates: "false",
		Boost:            "100",
		Stemming:         "false",
		Spelling:         "true",
		Logging:          "true",
		HtmlSettings: &searchblox.HtmlSettings{
			Description:    "test",
			MaxDocAge:      "1",
			MaxDocSize:     "100",
			SpiderMaxDepth: "1",
			SpiderMaxDelay: "1",
			UserAgent:      "httpcollection",
			Referer:        "Google",
			IgnoreRobots:   "false",
			FollowSitemap:  "false",
			FollowRedirect: "false",
		},
		BasicAuthSettings: &searchblox.BasicAuthSettings{
			Username: "searchblox",
			Password: "testing",
		},
		FormAuthSettings: &searchblox.FormAuthSettings{
			FormUrl:    "http://www.google.co.in",
			FormAction: "post",
			Form: []searchblox.Form{
				{
					Name:  "httpcollection",
					Value: "google",
				},
				{
					Name:  "httpcollection1",
					Value: "searchblox",
				},
			},
		},
		ProxySettings: &searchblox.ProxySettings{
			ServerUrl: "http://searchblox.com/proxy",
			Username:  "proxy",
			Password:  "adasd",
		},
	}
	_, err = c.UpdateSettingsInCollection(updateSettings)
	if err != nil {
		panic(err)
	}

	updateScheduler := searchblox.Collection{
		ApiKey:  apiKey,
		ColName: "httpcollection",
		Index: &searchblox.Scheduler{
			Frequency: "ONCE",
			TimeStamp: "21-01-2016 19:05:00",
		},
		Clear: &searchblox.Scheduler{
			Frequency: "MINUTELY",
			TimeStamp: "21-01-2016 18:05:00",
		},
		Refresh: &searchblox.Scheduler{
			Frequency: "WEEKLY",
			TimeStamp: "25-01-2016 30:05:00",
		},
	}
	_, err = c.UpdateSchedulerInCollection(updateScheduler)
	if err != nil {
		panic(err)
	}

	indexStopCrawler := searchblox.Collection{
		ApiKey:  apiKey,
		ColName: "httpcollection",
		Action:  "index",
	}
	_, err = c.IndexStopCrawlerInCollection(indexStopCrawler)
	if err != nil {
		panic(err)
	}
}
```
