package main

import (
	"fmt"
	"net/http"
	"products/database"
	"products/routes"
)

func main(){
	mux := http.NewServeMux()
	db, err := database.Init()
	if err != nil {
		fmt.Print("Erro ao iniciar o banco de dados")
	}

	routes.Init(mux, db)

	fmt.Print("Server Iniciado na Porta 3333\n")
	err = http.ListenAndServe("localhost:3333", mux)
	if err != nil {
		fmt.Print(err)
	}
}