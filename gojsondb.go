package gojsondb

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"log"

	"encoding/json"
)

// func main() {
// 	fmt.Println("Hello word")
// 	db, err := Load("user.json")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	usrs, err := db.Select("users")
	
// 	prettyPrint(usrs)

// }

func (gojsondb *GoJsonDb) Load() (interface{}, error) {

	var dbObject map[string]interface{}

	jsonFile, err := os.Open(gojsondb.DataPath)

	if err != nil {
		return nil, errors.New("File does not exist, please check file or file path")
	}

	defer jsonFile.Close()
	//Try unmarshal

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &dbObject)
	if err != nil {
		return nil, errors.New("Could not unmarshal json")
	}

	return dbObject, nil
}

func (gojsondb *GoJsonDb) Select(tableName string, data interface{}) (interface{}, error) {

	if data == nil {
		return nil, errors.New("The preloaded data is null")
	}


	v, ok := data.(map[string]interface{})
	if !ok {
		fmt.Println("not mappable")
	}

	if len(tableName) > 0 {
		return v[tableName], nil
	}else{
		return v, nil
	}
}

// func (jsonDB interface{}) Where(key string, value interface{}) (map[string]interface{}, error){
// 	if jsonDB == nil {
// 		return nil, errors.New("The preloaded data is null")
// 	}

// 	myMap := jsonDB.(map[string]interface{})

// 	for k, v := range jsonDB{
// 		fmt.Println("k:", k, "v:", v)
// 	}
// 	return nil, nil
// }

// func (jsonDB *JsonDB) Add(table string, data interface{}) (interface{}, error) {
// 	if jsonDB == nil {
// 		return nil, errors.New("The preloaded data is null")
// 	}

	
// }



func prettyPrint(i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func getProp(d interface{}, label string) (interface{}, bool) {
	switch reflect.TypeOf(d).Kind() {
	case reflect.Struct:
		v := reflect.ValueOf(d).FieldByName(label)
		
		return v.Interface(), true
	}
	return nil, false
}
