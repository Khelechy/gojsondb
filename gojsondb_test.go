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
		value, err := a.Select("plans", data)

		prettyPrint(value)
		
		if err != nil {
			t.Errorf("err: %v", err)
		}
	})
}
