package store

import (
	"encoding/json"
	. "fmt"
	"os"
)

type Store interface {
	Write(data interface{}) error
	Read(data interface{}) error
}

// Var Type of type string
type Type string

// Constant
const (
	FileType Type = "file"
)

// Struct with path of file
type FileStore struct {
	path string
}

// Write new data in the persistence file
func (fs *FileStore) Write(data interface{}) error {
	file, errJSON := json.Marshal(data)
	if errJSON != nil {
		return Errorf("error when make marshal in data: %s", errJSON)
	}
	return os.WriteFile(fs.path, file, 0644)
}

// Read data from persistence file
func (fs *FileStore) Read(data interface{}) error {
	file, errOpen := os.ReadFile(fs.path)
	if errOpen != nil {
		return Errorf("persistence file can not be open: %s", errOpen)
	}
	if errJSON := json.Unmarshal(file, &data); errJSON != nil {
		return Errorf("json data can not be unmarshal: %s", errJSON)
	}
	return nil
}

// Return a new Store Interface
func NewStore(store Type, pathFile string) Store {
	switch store {
	case FileType:
		return &FileStore{pathFile}
	}
	return nil
}
