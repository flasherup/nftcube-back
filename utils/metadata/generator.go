package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Attribute struct {
	ImageURL string `json:"image_url"`
	Title string `json:"title"`
	URL string `json:"url"`
}


func main() {
	size := 5329
	attributes := make(map[string]Attribute, size)
	for i := 0; i<size; i++ {
		key := strconv.FormatInt(int64(i+1), 10)
		attributes[key] = Attribute{
			ImageURL: "",
			Title: fmt.Sprintf("Tile#%s", key),
			URL: "",
		}
	}

	dataJson, _ := json.Marshal(attributes)
	err := ioutil.WriteFile("attributes.json", dataJson, 0644)
	if err != nil {
		fmt.Println("Attributes write error", err)
	}

	owners := make(map[string]string, size)
	for i := 0; i<size; i++ {
		key := strconv.FormatInt(int64(i+1), 10)
		owners[key] = "null"
	}

	dataJson, _ = json.Marshal(owners)
	err = ioutil.WriteFile("owners.json", dataJson, 0644)
	if err != nil {
		fmt.Println("Owners write error", err)
	}

}

