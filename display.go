package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Die(obj interface{}) {
	fmt.Println(obj)
	os.Exit(1)
}

func Display(obj interface{}) {
	str, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Println(string(str))
}
