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

/* TODO
To enable us to grab any value from Labels it should be a map.
https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/#unstructured-data-decoding-json-to-maps
*/
type Labels struct {
	App string `json:"app"`
}

// ReadDir reads the directory named by dirname and returns
// a list of directory entries
func ReadDir(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	/* I have no need for the sort but it's good to have for the future
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	*/
	var files []string

	for _, file := range list {
		files = append(files, file.Name())
	}

	return files, nil
}

// ReadItem marshalls the Struct Items and return a string to make it human readable.
func ReadItem(myitem Items) (string, error) {
	data, err := json.Marshal(myitem)
	// Notice the string convertion
	return string(data), err
}

func readjson(channel chan []byte, filename *string) {

	jsonFile, err := os.Open(*filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		/* TODO should i handle the print/log the errors here or should
		I have a channel where i send back the error.
		Use: https://godoc.org/golang.org/x/sync/errgroup
		*/
	}

	fmt.Println("Successfully Opened " + *filename)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	channel <- byteValue
}

// Parsejson stuff
func Parsejson(foldername string, myfiles []string) []Items {

	// Create the channel
	channel := make(chan []byte)
	defer close(channel)

	var listItems []Items
	for i := 0; i < len(myfiles); i++ {
		var items Items
		// TODO There is probably some nice path lib to make this look better.
		myfile := foldername + "/" + myfiles[i]

		go readjson(channel, &myfile)

		err := json.Unmarshal(<-channel, &items)
		if err != nil {
			fmt.Printf("Unable to parse %v, due to error: %v", myfile, err)
			// Using continue to jump out of the loop.
			continue
		}

		/*
			If you want to grab a specific value from the unmarshalled json.

			for i := 0; i < len(items.Items); i++ {
				fmt.Println("Item APIVersion " + items.Items[i].APIVersion)
				fmt.Println("Item Kind " + items.Items[i].Kind)
				fmt.Println("Item Metadata " + items.Items[i].Metadata.Name)
			}
		*/
		listItems = append(listItems, items)
	}
	return listItems
}
