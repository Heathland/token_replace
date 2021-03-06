package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var startToken, endToken string

var version, commit, date string

func main() {
	fileInput := flag.String("f", "", "file input")
	showHelp := flag.Bool("h", false, "show help")
	showVersion := flag.Bool("v", false, "version")
	flag.StringVar(&startToken, "startToken", "#{", "start token")
	flag.StringVar(&endToken, "endToken", "}#", "end token")
	flag.Parse()

	if *showVersion {
		fmt.Println("Token Replacer")
		fmt.Printf("Version: %v\n", version)
		fmt.Printf("Commit: %v\n", commit)
		fmt.Printf("Date: %v\n", date)
		os.Exit(0)
	}

	if *showHelp {
		fmt.Println("Token replacer syntax:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// If the file is being piped using "-f -" then we read from Pipe
	// Else we just take the path of the file being defined
	var output string
	var err error
	if *fileInput == "-" || *fileInput == "" {
		output, err = ReadFromPipe()
	} else {
		output, err = ReadFromFile(*fileInput)
	}

	if err != nil {
		panic(err)
	}

	// Output translated string
	fmt.Println(output)
}

func ReadFromFile(s string) (o string, err error) {
	data, err := os.ReadFile(s)
	if err != nil {
		return
	}
	o = RegexReplace(strings.TrimRight(string(data), "\n"))
	return
}

func ReadFromPipe() (o string, err error) {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		o += fmt.Sprintln(RegexReplace(reader.Text()))
	}

	if err = reader.Err(); err != nil {
		return
	}

	o = strings.TrimSuffix(o, "\n")
	return
}

func ReplaceWithEnv(s string) string {
	s = strings.ReplaceAll(s, startToken, "")
	s = strings.ReplaceAll(s, endToken, "")
	return os.Getenv(s)
}

func RegexReplace(s string) string {
	regex := regexp.MustCompile(fmt.Sprintf("%s([^}]+)%s", startToken, endToken))
	return regex.ReplaceAllStringFunc(s, ReplaceWithEnv)
}
