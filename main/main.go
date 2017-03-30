package main

import (
	"github.com/tokillamockingbird/searchblox-golang"
)

func main() {
	apiKey := "47AD645E72CEA9D8B2AB08A2312BF432"
	c := searchblox.Client{Host: "http://localhost:80"}

	// FIXME: check if we have limit on collection creation or error during creation situations
	// CREATE CUSTOM COLLECTION
	_, err := c.CreateCustomCollection(searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName: "Test",
		},
	})
	if err != nil {
		panic(err)
	}

	// INDEX DOCUMENT - CUSTOM COLLECTION
	_, err = c.IndexDocumentInCustomCollection(searchblox.Collection{
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
	})
	if err != nil {
		panic(err)
	}

	// DOCUMENT STATUS - CUSTOM COLLECTION
	_, err = c.DocumentStatusInCustomCollection(searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName: "Test",
			Uid:     "http://www.searchblox.com",
		},
	})
	if err != nil {
		panic(err)
	}

	// DELETE DOCUMENT - CUSTOM COLLECTION
	_, err = c.DeleteDocumentInCustomCollection(searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName: "Test",
			Uid:     "http://www.searchblox.com", // Maybe Location?
		},
	})
	if err != nil {
		panic(err)
	}

	// CLEAR CUSTOM COLLECTION
	_, err = c.ClearCustomCollection(searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName: "Test",
		},
	})
	if err != nil {
		panic(err)
	}

	// DELETE CUSTOM COLLECTION
	_, err = c.DeleteCustomCollection(searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName: "Test",
		},
	})
	if err != nil {
		panic(err)
	}

	// ADD/UPDATE DOCUMENT - HTTP/FILESYSTEM COLLECTION
	_, err = c.AddUpdateDocumentInCollection(searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName:  "Test",
			Location: "http://www.searchblox.com",
		},
	})
	if err != nil {
		panic(err)
	}

	// DELETE DOCUMENT - HTTP/FILESYSTEM COLLECTION
	_, err = c.DeleteDocumentInCollection(searchblox.Collection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName:  "Test",
			Location: "http://www.searchblox.com",
		},
	})
	if err != nil {
		panic(err)
	}

	// ADD HTTP, FILESYSTEM OR DATABASE COLLECTION
	_, err = c.AddCollection(searchblox.Collection{
		ApiKey:  apiKey,
		ColName: "httpcollection",
		ColType: "http",
	})
	if err != nil {
		panic(err)
	}

	// DELETE HTTP, FILESYSTEM OR DATABASE COLLECTION
	_, err = c.DeleteCollection(searchblox.Collection{
		ApiKey:  apiKey,
		ColName: "httpcollection",
	})
	if err != nil {
		panic(err)
	}

	// UPDATE PATH - HTTP, FILESYSTEM OR DATABASE COLLECTION
	_, err = c.UpdatePathInCollection(searchblox.Collection{
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
	})
	if err != nil {
		panic(err)
	}

	// UPDATE SETTINGS - HTTP, FILESYSTEM OR DATABASE COLLECTION
	_, err = c.UpdateSettingsInCollection(searchblox.Collection{
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
	})
	if err != nil {
		panic(err)
	}

	// UPDATE SCHEDULER - HTTP, FILESYSTEM OR DATABASE COLLECTION
	_, err = c.UpdateSchedulerInCollection(searchblox.Collection{
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
	})
	if err != nil {
		panic(err)
	}

	// INDEX/STOP CRAWLER - HTTP, FILESYSTEM OR DATABASE COLLECTION
	_, err = c.IndexStopCrawlerInCollection(searchblox.Collection{
		ApiKey:  apiKey,
		ColName: "httpcollection",
		Action:  "index",
	})
	if err != nil {
		panic(err)
	}
}
