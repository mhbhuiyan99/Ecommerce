package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id") // Extract productId from the request path

	pId, err := strconv.Atoi(productID) // Convert productId to integer if needed
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}

	product := database.Get(pId)

	if product == nil {
		util.SendError(w, 404, "Product not found")
		return
	}
	util.SendData(w, product, 200);
}