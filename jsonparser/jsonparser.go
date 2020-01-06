package jsonparser

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

/*`getValue` should be `GetValue` to be exposed to other packages.
It should start with a capital letter.
*/

func readjson(filename string) []byte {

	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println("Successfully Opened " + filename)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

// Parsejson stuff
func Parsejson() string {

	var items Items
	byteValue := readjson("small_svc.json")
	json.Unmarshal(byteValue, &items)

	for i := 0; i < len(items.Items); i++ {
		fmt.Println("Item APIVersion " + items.Items[i].APIVersion)
		fmt.Println("Item Kind " + items.Items[i].Kind)
		fmt.Println("Item Metadata " + items.Items[i].Metadata.Name)
	}

	return "hello"
}
