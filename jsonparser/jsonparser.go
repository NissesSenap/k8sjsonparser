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

// ReadDir reads the directory named by dirname and returns
// a list of directory entries
func readDir(dirname string) ([]string, error) {
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

func readjson(channel chan []byte, filename *string) {

	jsonFile, err := os.Open(*filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		/* TODO  how should i handle errors when i do channels?
		I don't want to do an exit here since i read multiple files.
		Feels like i should return them or is just a log the standard way?
		*/
		os.Exit(2)
	}

	fmt.Println("Successfully Opened " + *filename)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	channel <- byteValue
}

// Parsejson stuff
func Parsejson(foldername string) {

	myfiles, _ := readDir(foldername)

	// Create the channel
	channel := make(chan []byte)
	defer close(channel)

	for i := 0; i < len(myfiles); i++ {
		var items Items
		// TODO There is probably some nice path lib to make this look better.
		myfile := foldername + "/" + myfiles[i]

		go readjson(channel, &myfile)

		json.Unmarshal(<-channel, &items)

		for i := 0; i < len(items.Items); i++ {
			fmt.Println("Item APIVersion " + items.Items[i].APIVersion)
			fmt.Println("Item Kind " + items.Items[i].Kind)
			fmt.Println("Item Metadata " + items.Items[i].Metadata.Name)
		}
	}
}
