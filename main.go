package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"encoding/json"
)

type JsonDB struct {
	jsonObject map[string]interface{}
}

func main() {
	fmt.Println("Hello word")
	db, err := Load("user.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(db.jsonObject["users"])
}

func Load(filePathLoaded string) (JsonDB, error) {
	//data, err := ioutil.ReadFile(filePathLoaded)
	jsonDB := JsonDB{}
	var dbObject map[string]interface{}

	jsonFile, err := os.Open(filePathLoaded)

	if err != nil {
		return jsonDB, errors.New("File does not exist, please check file or file path")
	}

	defer jsonFile.Close()
	//Try unmarshal

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &dbObject)
	if err != nil {
		return jsonDB, errors.New("Could not unmarshal json")
	}

	returnObject := &JsonDB{
		jsonObject: dbObject,
	}

	return *returnObject, nil
}

func (jsonDB *JsonDB) Select(table string) interface{} {
	return jsonDB.jsonObject[table]
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func getProp(d interface{}, label string) (interface{}, bool) {
	switch reflect.TypeOf(d).Kind() {
	case reflect.Struct:
		v := reflect.ValueOf(d).FieldByName(label)
		return v.Interface(), true
	}
	return nil, false
}
