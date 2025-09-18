package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id") // Extract productId from the request path

	pId, err := strconv.Atoi(productID) // Convert productId to integer if needed
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.productRepo.Get(pId)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}
	util.SendData(w, http.StatusOK, product);
}