package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	bytes, err := json.Marshal(colors)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

