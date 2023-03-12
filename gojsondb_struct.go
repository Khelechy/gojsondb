package gojsondb

type GoJsonDb struct {
	DataPath string
	Functions
}

type Functions interface {
	Load() (interface{}, error)
	Select(data interface{}, tableName string) (interface{}, error)
	Where(data interface{}, key string, value interface{}) (interface{}, error)
	Add(data interface{}, tableName string, newData interface{}) error
}
