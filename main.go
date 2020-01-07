package main

import (
	"github.com/NissesSenap/k8sjsonparser/jsonparser"
)

func main() {
	/* Scan the folder that is defined in Parsejson
	returns a list.
	Send that list to be read by readjson
	Unmarshall the result and print it all.
	*/

	jsonparser.Parsejson("jsonfiles")
}
