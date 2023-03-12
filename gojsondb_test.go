package gojsondb

import "testing"

func TestLoad(t *testing.T){
	t.Run("Run load json", func(t *testing.T) {
		a := GoJsonDb{
			DataPath: "user.json",
		}

		data, err := a.Load()

		prettyPrint(data)
		
		if err != nil {
			t.Errorf("err: %v", err)
		}
	})
}

func TestSelect(t *testing.T){
	t.Run("Run select table", func(t *testing.T) {
		a := GoJsonDb{
			DataPath: "user.json",
		}

		data, err := a.Load()
		value, err := a.Select(data, "users")

		prettyPrint(value)
		
		if err != nil {
			t.Errorf("err: %v", err)
		}
	})
}

func TestWhere(t *testing.T){
	t.Run("Run where", func(t *testing.T) {
		a := GoJsonDb{
			DataPath: "user.json",
		}

		data, err := a.Load()
		value, err := a.Select(data, "users")
		newValue, err := a.Where(value, "verified", true)

		prettyPrint(newValue)
		
		if err != nil {
			t.Errorf("err: %v", err)
		}
	})
}

func TestAdd(t *testing.T){
	t.Run("Run where", func(t *testing.T) {
		a := GoJsonDb{
			DataPath: "user.json",
		}

		data, err := a.Load()

		person := struct {
			id int
			firstname string
			lastname string
			age  int
			verified bool
		}{
			id: 4,
			firstname: "jane",
			lastname: "mark",
			age:  30,
			verified: false,
		}
		
		var v interface{} = person

		// var v interface{} = map[string]interface{}{
		// 	"id": 4,
		// 	"firstname": "jane",
		// 	"lastname": "mark",
		// 	"age":  30,
		// 	"verified": false,
		// }

		err = a.Add(data, "users", v)
		
		if err != nil {
			t.Errorf("err: %v", err)
		}
	})
}
