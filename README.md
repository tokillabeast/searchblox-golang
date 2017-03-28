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
    docker-compose up -d
```
After this you can check Searchblox on 8089 port, open `localhost:8089` in browser to check it.

```go
package main

import (
	"github.com/tokillamockingbird/searchblox-golang/client"
)

func main() {
	c := client.SearchBlox{Host: "http://localhost:8089"}
	customCollection := client.CustomCollection{
		ApiKey: "25B213BA03FAB750790FC63FD1C6B301",
		Document: client.Document{
			Colname: "Test",
		},
	}
	err := c.CreateCustomCollection(customCollection)
	if err != nil {
		panic(err)
	}
	indexCustomCollection := client.CustomCollection{
		ApiKey: "25B213BA03FAB750790FC63FD1C6B301",
		Document: client.Document{
			Colname: "Test",
			Url : "http://www.searchblox.com",
			Uid : "http://www.searchblox.com",
			Location:"http://www.searchblox.com",
			Alpha : "string",
			Size : "44244",
			Title : "Text",
			Keywords : "keywords",
			Description : "SearchBlox Content Search Software",
			Content : "content",
			LastModified : "14 January 2015 06:19:42 GMT",
			ContentType : "HTML",
			Meta: client.Meta {
				Location: "San Francisco",
				Temp: "23",
				Weather: "sunny",
			},
		},
	}
	err = c.IndexDocumentCustomCollection(indexCustomCollection)
	if err != nil {
		panic(err)
	}
	err = c.ClearCustomCollection(customCollection)
	if err != nil {
		panic(err)
	}
	err = c.DeleteCustomCollection(customCollection)
	if err != nil {
		panic(err)
	}
}
```
