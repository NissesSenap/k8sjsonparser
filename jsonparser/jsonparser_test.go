package jsonparser

import (
	"testing"
)

func TestReadDir(t *testing.T) {
	/* Make test better don't want to fail if i add a file that have a file
	with a letter earlier then s
	*/
	myfiles, _ := ReadDir("testdata")
	if myfiles[0] != "small_svc.json" {
		t.Errorf("File is not in folder %v", myfiles[0])
	}
}
