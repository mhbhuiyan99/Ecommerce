package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImageURL string `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id") // Extract productId from the request path

	pId, err := strconv.Atoi(productID) // Convert productId to integer if needed
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	
	var req ReqUpdateProduct
	
	decoder := json.NewDecoder(r.Body) // take the information from frontend 
	err = decoder.Decode(&req) // decoder decodes the JSON data from r.Body into the newProduct variable and place it in the address of newProduct (&newProduct)

	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	_, err = h.productRepo.Update(repo.Product{
		ID:	pId, 
		Title: req.Title, 
		Description:  req.Description, 
		Price: req.Price, 
		ImageURL: req.ImageURL,
	})
	if err != nil {	
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	util.SendData(w, http.StatusOK ,"Successfully updated product")
}