package review

import (
	"net/http"
	"encoding/json"
	"fmt"
	"ecommerce/database"
	"ecommerce/util"
)

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {

	var newUser database.User
	
	decoder := json.NewDecoder(r.Body) // take the information from frontend 
	err := decoder.Decode(&newUser) // decoder decodes the JSON data from r.Body into the newProduct variable and place it in the address of newProduct (&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}
	createUser := newUser.Store()

	util.SendData(w, createUser, http.StatusCreated) 
}