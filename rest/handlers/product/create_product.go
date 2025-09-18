package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImageURL string `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateProduct
	
	decoder := json.NewDecoder(r.Body) // take the information from frontend 
	err := decoder.Decode(&req) // decoder decodes the JSON data from r.Body into the newProduct variable and place it in the address of newProduct (&newProduct)

	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	createProduct, err := h.productRepo.Create(repo.Product{
		Title: req.Title,
		Description: req.Description,
		Price: req.Price,
		ImageURL: req.ImageURL,
	})

	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, http.StatusCreated, createProduct) 
}

