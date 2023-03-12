package gojsondb

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	//"strconv"

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

type jsonarray struct {
    data []interface{}
}


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

func (gojsondb *GoJsonDb) Select(data interface{}, tableName string, ) (interface{}, error) {

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

func (gojsondb *GoJsonDb) Where(data interface{}, key string, value interface{}) (interface{}, error) {

	if data != nil {
		newData, err := json.Marshal(data)
		if err != nil {
			fmt.Print("Can not unmarshal")
		}		
        dataArray, err := parseData(newData)
        var newDataArray []interface{}
        for _, singleData := range dataArray.data {
            dataMap := singleData.(map[string]interface{})
			switch v := dataMap[key].(type){
			case float64:
				if int(v) == value {
					newDataArray = append(newDataArray, singleData)
				}
			case string:
				if v == value {
					newDataArray = append(newDataArray, singleData)
				}
			case bool:
				if bool(v) == value {
					newDataArray = append(newDataArray, singleData)
				}
			}
        }
        return newDataArray, nil
    } else {
        return nil, errors.New("The preloaded data is null")
    }

}


func (gojsondb *GoJsonDb) Add(data interface{}, tableName string, newData interface{}) error {
	
	if data != nil {

		v, ok := data.(map[string]interface{})
		if !ok {
			fmt.Println("not mappable")
		}

		var selectedData interface{}

		if len(tableName) > 0 {
			selectedData = v[tableName]
		}else{
			selectedData = v
		}

		dataBytes, err := json.Marshal(selectedData)
		if err != nil {
			fmt.Print("Can not unmarshal")
		}		
        dataArray, err := parseData(dataBytes)
		if err != nil {
			fmt.Println("Can not parse data")
		}

		//Transform New Data 
		var finalData interface{}
		newDataMap, ok := newData.(map[string]interface{})
		if !ok {
			fmt.Println("not mappable transform")
			finalData = newDataMap
		}else{
			finalData = newData
		}

		//prettyPrint(newDataMap)

		dataArray.data = append(dataArray.data, finalData)
		writeData , err := json.Marshal(dataArray.data)
		if err != nil {
			fmt.Println("can not marshal write data")
		}

    	err = ioutil.WriteFile(gojsondb.DataPath, writeData, 0644)

		return nil
	}else {
        return errors.New("The preloaded data is null")
    }

	
}



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

func parseData(data []byte) (dataArray *jsonarray, err error) {
    if data != nil {
        var token interface{}
        err = json.Unmarshal(data, &token)
        if err != nil {
            return nil, err
        }
        switch token.(type) {
        case []interface{}:
            dataArray = &jsonarray{token.([]interface{})}
            return dataArray, nil
        default:
            return nil, errors.New("The preloaded data is not a Json Array")
        }
    } else {
        return nil, errors.New("The preloaded data is null")
    }
}
