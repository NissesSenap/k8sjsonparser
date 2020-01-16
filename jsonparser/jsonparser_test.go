package jsonparser

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

type IndexDirTest struct {
	testdir     string
	expectedOut []string
}

var myIndex = []IndexDirTest{
	IndexDirTest{"testdata", []string{"small_svc.json"}},
	IndexDirTest{"ThisDirDontExist", []string{""}},
}

func TestReadDirReal(t *testing.T) {
	// The second test works since err handels it and dosen't create any actual since it's Nil
	for _, tests := range myIndex {
		actualOutput, _ := ReadDir(tests.testdir)
		for i := 0; i < len(actualOutput); i++ {
			if actualOutput[i] != tests.expectedOut[i] {
				t.Errorf("The expected files from: %v didn't matche: %v from folder: %v", tests.expectedOut[i], actualOutput[i], tests.testdir)
			}
		}
	}
}

func TestReadDirError(t *testing.T) {
	// Test the error case to make sure we get an error.
	_, err := ReadDir("FolderDontExist")
	if err == nil {
		t.Errorf("Crap %v", err)
	}
}

func TestReadJson(t *testing.T) {
	// var expectedByte []byte
	myfile := "testdata/small_svc.json"

	// Create channel
	channel := make(chan []byte)
	defer close(channel)

	// Send data to readjson
	go readjson(channel, &myfile)
	testByte := <-channel

	// Read the file in the test
	jsonFile, _ := os.Open(myfile)
	defer jsonFile.Close()
	expectedByte, _ := ioutil.ReadAll(jsonFile)

	// Compare the two []bytes if 0 it's the same data
	lala := bytes.Compare(testByte, expectedByte)

	if lala != 0 {
		t.Errorf("Bytes is not the same: %v", myfile)
	}
}
