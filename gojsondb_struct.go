package gojsondb

type GoJsonDb struct {
	DataPath string
	Functions
}

type Functions interface {
	Load() (interface{}, error)
	Select(tableName string, data interface{}) (interface{}, error)
}
