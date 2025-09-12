package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	
	decoder := json.NewDecoder(r.Body) // take the information from frontend 
	err := decoder.Decode(&newProduct) // decoder decodes the JSON data from r.Body into the newProduct variable and place it in the address of newProduct (&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "please give me valid JSON", 400)
		return
	}

	createProduct := database.Store(newProduct)

	util.SendData(w, createProduct, 201) 
}

