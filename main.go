package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Items struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Items      []Item `json:"items"`
}

type Item struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
}

type Metadata struct {
	Labels    Labels `json:"labels"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type Labels struct {
	App string `json:"app"`
}

func main() {
	jsonFile, err := os.Open("small_svc.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println("Successfully Opened small_svc.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var items Items

	json.Unmarshal(byteValue, &items)

	for i := 0; i < len(items.Items); i++ {
		fmt.Println("Item APIVersion " + items.Items[i].APIVersion)
		fmt.Println("Item Kind " + items.Items[i].Kind)
		fmt.Println("Item Metadata " + items.Items[i].Metadata.Name)
	}
}
