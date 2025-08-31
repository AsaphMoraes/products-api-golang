package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"products/database"
)

func HandleGetAllProducts(p Params) {
	produtos, err := p.db.GetAllProduct()
	if err != nil{
		fmt.Print(err)
	}

	b, err := json.Marshal(produtos)
	if err != nil{
		fmt.Print(err)
	}
	
	p.res.Write(b)
}

//Criar arquivos separados para lidar com cada func separadamente

func HandleGetProductByBrand(w http.ResponseWriter, r *http.Request, db *database.Database) {
	produtos, err := db.GetAllProduct()
	if err != nil{
		fmt.Print(err)
	}

	b, err := json.Marshal(produtos)
	if err != nil{
		fmt.Print(err)
	}
	
	w.Write(b)
}

func HandleGetProductByInInventory(w http.ResponseWriter, r *http.Request, db *database.Database) {
	produtos, err := db.GetAllProduct()
	if err != nil{
		fmt.Print(err)
	}

	b, err := json.Marshal(produtos)
	if err != nil{
		fmt.Print(err)
	}
	
	w.Write(b)
}