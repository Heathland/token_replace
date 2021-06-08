package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	_, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var regex = regexp.MustCompile("#{([^}]+)}#")
		s := regex.ReplaceAllStringFunc(scanner.Text(), ReplaceWithEnv)
		fmt.Println(s)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func ReplaceWithEnv(s string) string {
	s = strings.ReplaceAll(s, "{", "")
	s = strings.ReplaceAll(s, "}", "")
	s = strings.ReplaceAll(s, "#", "")
	return os.Getenv(s)
}
