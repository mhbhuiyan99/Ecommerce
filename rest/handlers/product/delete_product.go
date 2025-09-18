package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id") // Extract productId from the request path

	pId, err := strconv.Atoi(productID) // Convert productId to integer if needed
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	err = h.productRepo.Delete(pId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, http.StatusOK, "Successfully deleted product")
}