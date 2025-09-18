package product

import (
	"net/http"
	"ecommerce/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) { // r information about resoures, return requested resource(products) by w)
	productList, err := h.productRepo.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	util.SendData(w, http.StatusOK, productList) 
}