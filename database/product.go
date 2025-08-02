package database

import (
	"encoding/json"
	"fmt"
	"os"
	"products/models"
)

type FunctionsProduct interface{
	RegisterProduct(models.Product) error
	UpdatePrice(id uint8, newValue float32) error
	UpdateInInventory(id uint8, newValue uint8) error
	GetProductById(id uint8) (models.Product, error)
	GetProductByBrand(brand string) ([]models.Product, error)
	GetAllProduct() ([]models.Product, error)
	GetProductInInventory(min uint8, max uint8) ([]models.Product, error)
}

func read(path string)[]models.Product{
	b, err := os.ReadFile(path)
	if err != nil{
		fmt.Print("Erro ao acessar o banco de dados")
	}
	data := []models.Product{}
	json.Unmarshal(b, &data)

	return data
}

func write(path string, p []models.Product)error{
	b, err := json.Marshal(p)
	if err != nil{
		return err
	}

	os.WriteFile(path, b, 0666)

	return err
}

func (d *Database) RegisterProduct(p models.Product)error{
	data := read(d.path)
	newData := append(data, p)

	err := write(d.path, newData)
	if err != nil{
		fmt.Print("Erro ao registrar o produto")
	}

	return err
}

func (d *Database) UpdatePrice(id uint8, newValue float32)error{
	data := read(d.path)

	for i, p := range data{
		if p.Id == id{
			data[i].Price = newValue
		}
	}

	err := write(d.path, data)
	if err != nil{
		fmt.Print("Erro ao registrar o produto")
	}

	return err
}

func (d *Database) UpdateInInventory(id uint8, newValue uint8)error{
	data := read(d.path)

	for i, p := range data{
		if p.Id == id{
			data[i].InInventory = newValue
		}
	}

	err := write(d.path, data)
	if err != nil{
		fmt.Print("Erro ao registrar o produto")
	}

	return err
}

func (d *Database) GetProductById(id uint8) (models.Product, error){
	data := read(d.path)

	if id > uint8(len(data)) {
		return models.Product{}, fmt.Errorf("ID nao existe")
	}

	for _, p := range data{
		if p.Id == id {
			return p, nil
		}
	}

	return models.Product{}, fmt.Errorf("produto nao encontrado")
}

func (d *Database) GetProductByBrand(brand string) ([]models.Product, error){
	data := read(d.path)
	products := []models.Product{}

	for _, p := range data{
		if p.Brand == brand {
			products = append(products, p)
		}
	}
	if len(products) == 0 {
		return products, fmt.Errorf("nenhum produto possui a marca: %v", brand)
	}

	return products, nil
}

func (d *Database) GetAllProduct()([]models.Product, error){
	data := read(d.path)

	if len(data) == 1{
		return []models.Product{}, fmt.Errorf("nao possui produtos")
	}

	return data, nil
}

func (d *Database) GetProductInInventory(min uint8, max uint8)([]models.Product, error){
	data := read(d.path)
	products := []models.Product{}

	for _, p := range data{
		if p.InInventory > min || p.InInventory < max{
			products = append(products, p)
		}
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("produto nao encontrado")
	}

	return products, nil
}
