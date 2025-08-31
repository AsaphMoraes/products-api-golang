package routes

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func HandleGetProductById(p Params) {
	id, err := strconv.ParseUint(p.data[0], 0, 8)
	if err != nil {
		fmt.Print(err)
	}
	
	//Modificar a saida da api para quando o produto n existir, ao inves de um json incorreto, modificar para uma msg onde informe q n existe o produto.
	produtos, err := p.db.GetProductById(uint8(id))
	if err != nil{
		fmt.Print(err)
	}

	b, err := json.Marshal(produtos)
	if err != nil{
		fmt.Print(err)
	}
	
	p.res.Write(b)
}