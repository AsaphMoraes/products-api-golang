package routes

import (
	"net/http"
	"products/database"
)

type Params struct{
	res http.ResponseWriter
	req *http.Request
	db *database.Database
	data []string
}

func Init(mux *http.ServeMux, db *database.Database){
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		query := r.URL.Query()

		id := query.Get("id")
		brand := query.Get("brand")
		min := query.Get("min")
		max := query.Get("max")

		//adicionar correspondentemente os outros metodos da api.
		if r.Method == http.MethodGet{
			if id != "" || brand == "" || min == "" || max == "" {
				HandleGetProductById(Params{
					res: w,
					req: r,
					db: db,
					data: []string{id},
				})
				return
			}

			HandleGetAllProducts(Params{
					res: w,
					req: r,
					db: db,
					data: []string{},
				})

			return
		}

		http.Error(w, "", http.StatusMethodNotAllowed)
	})
}