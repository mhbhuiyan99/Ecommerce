package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id") // Extract productId from the request path

	pId, err := strconv.Atoi(productID) // Convert productId to integer if needed
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}
	
	var newProduct database.Product
	
	decoder := json.NewDecoder(r.Body) // take the information from frontend 
	err = decoder.Decode(&newProduct) // decoder decodes the JSON data from r.Body into the newProduct variable and place it in the address of newProduct (&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "please give me valid JSON", 400)
		return
	}

	newProduct.ID = pId
	database.Update(newProduct)
	util.SendData(w, "Successfully updated product", 201)
}