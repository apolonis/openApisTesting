package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Name      string      `json:"name"`
	Countries []Countries `json:"data"`
}

// A Pokemon Struct to map every pokemon to.
type Countries struct {
	Country string `json:"country`
}

func main() {
	response, err := http.Get("https://api.coingecko.com/api/v3/events/countries")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(responseData))
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Countries))

	for i := 0; i < len(responseObject.Countries); i++ {
		fmt.Println(responseObject.Countries[i].Country)
	}
}
