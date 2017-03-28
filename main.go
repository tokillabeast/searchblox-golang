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
	err = c.ClearCustomCollection(customCollection)
	if err != nil {
		panic(err)
	}
	err = c.DeleteCustomCollection(customCollection)
	if err != nil {
		panic(err)
	}
}
