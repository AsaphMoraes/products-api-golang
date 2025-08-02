package routes

import (
	"net/http"
)

func Init(mux *http.ServeMux){
	mux.HandleFunc("/products", func(http.ResponseWriter, *http.Request){})
}