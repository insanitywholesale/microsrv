package handlers

import (
	"microsrv/apisrv/data"
	"net/http"
)

// swagger:route POST /products products addProduct
// Create a new product
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

//Create handles POST requests to add new products
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Printf("[DEBUG]: Handle POST")
	//Fetch the product from the request context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("[DEBUG]: Inserting product: %#v\n", prod)
	data.AddProduct(prod)
}

//Old method for adding a new product
//func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
//	p.l.Println("Handle POST Product")
//
//	prod := r.Context().Value(KeyProduct{}).(data.Product)
//	data.AddProduct(prod)
//}
