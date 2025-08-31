package database

import (
	"fmt"
	"os"
	"path"
)

func Init() (*Database, error){
	dirname, err := os.Getwd()
	if err != nil{
		fmt.Print(err)
	}
	databasePath := path.Join(dirname,"database", "products.json")

	file, err := os.Open(databasePath)
	if os.IsNotExist(err){
		_, err = os.Create(databasePath)
		if err != nil{
			return nil, fmt.Errorf("não foi possivel criar o banco de dados")
		}
	}
	defer file.Close()

	return &Database{Path: databasePath}, nil
}
