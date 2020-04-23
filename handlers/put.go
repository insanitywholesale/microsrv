package handlers

import (
	"microsrv/data"
	"net/http"
)

// swagger:route PUT /products products updateProduct
// Replaces one product
// responses:
//  201: noContentResponse
//  404: errorResponse
//  422: errorValidation
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] updating record id", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrorProductNotFound {
		p.l.Println("[ERROR] product not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

	//write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
