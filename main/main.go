package main

import (
	"github.com/tokillamockingbird/searchblox-golang"
)

func main() {
	apiKey := "47AD645E72CEA9D8B2AB08A2312BF432"
	c := searchblox.Client{Host: "http://localhost:80"}

	customCollection := searchblox.CustomCollection{
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

	indexCustomCollection := searchblox.CustomCollection{
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
	_, err = c.IndexDocumentCustomCollection(indexCustomCollection)
	if err != nil {
		panic(err)
	}

	statusCustomCollection := searchblox.CustomCollection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName: "Test",
			Uid:     "http://www.searchblox.com",
		},
	}
	_, err = c.DocumentStatusCustomCollection(statusCustomCollection)
	if err != nil {
		panic(err)
	}

	deleteCustomCollection := searchblox.CustomCollection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName: "Test",
			Uid:     "http://www.searchblox.com", // Maybe Location?
		},
	}
	_, err = c.DeleteDocumentCustomCollection(deleteCustomCollection)
	if err != nil {
		panic(err)
	}

	addUpdateDocumentHttpFileSystemCollection := searchblox.CustomCollection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName:  "Test",
			Location: "http://www.searchblox.com",
		},
	}
	_, err = c.AddUpdateDocumentHttpFileSystemCollection(addUpdateDocumentHttpFileSystemCollection)
	if err != nil {
		panic(err)
	}

	deleteDocumentHttpFileSystemCollection := searchblox.CustomCollection{
		ApiKey: apiKey,
		Document: &searchblox.Document{
			ColName:  "Test",
			Location: "http://www.searchblox.com",
		},
	}
	_, err = c.DeleteDocumentHttpFileSystemCollection(deleteDocumentHttpFileSystemCollection)
	if err != nil {
		panic(err)
	}

	addHttpFileSystemOrDatabaseCollection := searchblox.CustomCollection{
		ApiKey:  apiKey,
		ColName: "httpcollection",
		ColType: "http",
	}
	_, err = c.AddHttpFileSystemOrDatabaseCollection(addHttpFileSystemOrDatabaseCollection)
	if err != nil {
		panic(err)
	}

	deleteHttpFileSystemOrDatabaseCollection := searchblox.CustomCollection{
		ApiKey:  apiKey,
		ColName: "httpcollection",
	}
	_, err = c.DeleteHttpFileSystemOrDatabaseCollection(deleteHttpFileSystemOrDatabaseCollection)
	if err != nil {
		panic(err)
	}

	updatePathHttpFileSystemOrDatabaseCollection := searchblox.CustomCollection{
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
	_, err = c.UpdatePathHttpFileSystemOrDatabaseCollection(updatePathHttpFileSystemOrDatabaseCollection)
	if err != nil {
		panic(err)
	}

	updateSettingsHttpFileSystemOrDatabaseCollection := searchblox.CustomCollection{
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
	_, err = c.UpdateSettingsHttpFileSystemOrDatabaseCollection(updateSettingsHttpFileSystemOrDatabaseCollection)
	if err != nil {
		panic(err)
	}

	updateSchedulerHttpFileSystemOrDatabaseCollection := searchblox.CustomCollection{
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
	_, err = c.UpdateSchedulerHttpFileSystemOrDatabaseCollection(updateSchedulerHttpFileSystemOrDatabaseCollection)
	if err != nil {
		panic(err)
	}

	indexStopCrawlerHttpFileSystemOrDatabaseCollection := searchblox.CustomCollection{
		ApiKey:  apiKey,
		ColName: "httpcollection",
		Action:  "index",
	}
	_, err = c.IndexStopCrawlerHttpFileSystemOrDatabaseCollection(indexStopCrawlerHttpFileSystemOrDatabaseCollection)
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
}
