package main

import (
	"net/http"
	"products/routes"
)

func main(){
	mux := http.NewServeMux()
	
	routes.Init(mux)
}