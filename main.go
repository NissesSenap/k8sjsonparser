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
	jsonparser.Parsejson(foldername, myfiles)

}
