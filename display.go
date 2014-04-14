package main

import (
	"encoding/json"
	"fmt"
)

func Display(obj interface{}) {
	str, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Println(string(str))
}
