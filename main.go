package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"encoding/json"
)

var dbObject interface{}

func main() {
	fmt.Println("Hello word")
	Load("user.json")
}

func Load(filePathLoaded string) (interface{}, error) {
	data, err := ioutil.ReadFile(filePathLoaded)
	if err != nil {
		return nil, errors.New("File does not exist, please check file or file path")
	}
	//Try unmarshal

	err = json.Unmarshal(data, &dbObject)
	if err != nil {
		return nil, errors.New("Could not unmarshal json")
	}

	return &dbObject, nil
}
