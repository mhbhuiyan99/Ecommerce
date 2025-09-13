package product

import (
	"net/http"
	"ecommerce/database"
	"ecommerce/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) { // r information about resoures, return requested resource(products) by w)
	util.SendData(w, database.List(), 200) 
}