package main

import (
	"fmt"
	"os"
	"strings"
)

func StringToDict(line string) map[string]string {
	dict := make(map[string]string)

	split := strings.Split(line, ",")
	for _, value := range split {
		split2 := strings.Split(value, "=")
		if len(split2) == 2 {
			dict[split2[0]] = split2[1]
		}
	}

	return dict
}

func StringToList(line string) []string {
	return strings.Split(line, ",")
}

func Die(obj interface{}) {
	fmt.Println(obj)
	os.Exit(1)
}
