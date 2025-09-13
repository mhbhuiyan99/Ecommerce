package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id") // Extract productId from the request path

	pId, err := strconv.Atoi(productID) // Convert productId to integer if needed
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}
	database.Delete(pId)
	util.SendData(w, "Successfully deleted product", 201)
}