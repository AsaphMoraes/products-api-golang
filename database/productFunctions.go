package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type BuilderProduct interface{
	RegisterProduct(Product) error
	UpdatePrice(id uint8, newValue float32) error
	UpdateInInventory(id uint8, newValue uint8) error
	GetProductById(id uint8) (Product, error)
	GetProductByBrand(brand string) ([]Product, error)
	GetAllProduct() ([]Product, error)
	GetProductInInventory(min uint8, max uint8) ([]Product, error)
}

func read(path string) (data []Product){
	b, err := os.ReadFile(path)
	if err != nil{
		fmt.Print("Erro ao acessar o banco de dados")
	}

	json.Unmarshal(b, &data)

	return data
}

func write(path string, p []Product)error{
	b, err := json.Marshal(p)
	if err != nil{
		return err
	}

	os.WriteFile(path, b, 0666)

	return err
}

func (d *Database) RegisterProduct(p Product)error{
	data := read(d.Path)
	newData := append(data, p)

	err := write(d.Path, newData)
	if err != nil{
		fmt.Print("Erro ao registrar o produto")
	}

	return err
}

func (d *Database) UpdatePrice(id uint8, newValue float32)error{
	data := read(d.Path)

	for i, p := range data{
		if p.Id == id{
			data[i].Price = newValue
		}
	}

	err := write(d.Path, data)
	if err != nil{
		fmt.Print("Erro ao registrar o produto")
	}

	return err
}

func (d *Database) UpdateInInventory(id uint8, newValue uint8)error{
	data := read(d.Path)

	for i, p := range data{
		if p.Id == id{
			data[i].InInventory = newValue
		}
	}

	err := write(d.Path, data)
	if err != nil{
		fmt.Print("Erro ao registrar o produto")
	}

	return err
}

func (d *Database) GetProductById(id uint8) (Product, error){
	data := read(d.Path)

	//modificar logica para informar q o id n existe.
	if id > uint8(len(data)) {
		return Product{}, fmt.Errorf("ID nao existe")
	}

	for _, p := range data{
		if p.Id == id {
			return p, nil
		}
	}

	return Product{}, fmt.Errorf("produto nao encontrado")
}

func (d *Database) GetProductByBrand(brand string) ([]Product, error){
	data := read(d.Path)
	products := []Product{}

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

func (d *Database) GetAllProduct()([]Product, error){
	data := read(d.Path)

	if len(data) < 1{
		return []Product{}, fmt.Errorf("nao possui produtos")
	}

	return data, nil
}

func (d *Database) GetProductInInventory(min uint8, max uint8)([]Product, error){
	data := read(d.Path)
	products := []Product{}

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
