package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	err := os.Setenv("TEST", "foobar")

	startToken = "#{"
	endToken = "}#"
	testFile := "examples/standard.yaml"
	want := "value: foobar"


	testOutput, err := ReadFromFile(testFile)
	if testOutput != want || err != nil {
		t.Fatalf(`Invalid response when translating file. Have: "%s" Want: "%s" Error: "%v"`, testOutput, want, err)
	}
}

func TestReadFromPipe(t *testing.T) {
	err := os.Setenv("TEST", "foobar")

	startToken = "#{"
	endToken = "}#"
	testInput := "value: #{TEST}#"
	want := "value: foobar"

	content := []byte(testInput)
	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpFile.Name()) // clean up

	if _, err := tmpFile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpFile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpFile

	testOutput, err := ReadFromPipe()
	if testOutput != want || err != nil {
		t.Fatalf(`Invalid response when translating file. Have: "%s" Want: "%s" Error: "%v"`, testOutput, want, err)
	}
}

func TestCustomTokenReadFromFile(t *testing.T) {
	err := os.Setenv("TEST", "foobar")

	startToken = "_+"
	endToken = "+_"
	testFile := "examples/custom.yaml"
	want := "value: foobar"


	testOutput, err := ReadFromFile(testFile)
	if testOutput != want || err != nil {
		t.Fatalf(`Invalid response when translating file. Have: "%s" Want: "%s" Error: "%v"`, testOutput, want, err)
	}
}