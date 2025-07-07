package files

import (
	"demo/app-1/output"
	"fmt"
	"os"
)

type JsonDB struct {
	filename string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		filename: name,
	}
}

func (db *JsonDB) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		output.PrintError(err)
	}

	fmt.Println("Запись успешна")
}

func (db *JsonDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		output.PrintError(err)
		return nil, err
	}

	return data, nil
}
