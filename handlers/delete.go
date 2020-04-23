package handlers

import (
	"microsrv/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrorProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

//Old method to delete a product
//func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
//	p.l.Println("Handle DELETE Product")
//
//	muxVars := mux.Vars(r)
//	id, err := strconv.Atoi(muxVars["id"])
//	if err != nil {
//		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
//		return
//	}
//
//	err = data.DeleteProduct(id)
//	if err == data.ErrorProductNotFound {
//		http.Error(rw, "Product not found", http.StatusNotFound)
//		return
//	}
//	if err != nil {
//		http.Error(rw, "Product not found", http.StatusInternalServerError)
//		return
//	}
//}
