package main

import (
	"fmt"
	"os"

	"github.com/NissesSenap/k8sjsonparser/jsonparser"
)

func main() {
	/* Scan the folder that is defined in Parsejson
	returns a list.
	Send that list to be read by readjson
	Unmarshall the result and print it all.
	*/

	foldername := "jsonfiles"
	myfiles, err := jsonparser.ReadDir(foldername)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	myitems := jsonparser.Parsejson(foldername, myfiles)

	// Print out the json files in a human readable format.
	for i := 0; i < len(myitems); i++ {
		data, err := jsonparser.ReadItem(myitems[i])
		if err != nil {
			fmt.Printf("Soemthing wen't wrong creating the json string %v", err)
		}
		fmt.Println(data)
	}
}
