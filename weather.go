package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type forecast struct {
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please provide a location for a weather forecast")
	} else {
		location = os.Args[1]
		fmt.Printf("Looking up: %s\n", location)
	}
}
