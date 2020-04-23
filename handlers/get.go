package handlers

import (
	"microsrv/data"
	"net/http"
)

// swagger:route GET /products products getProducts
// Return a list of products from the data store
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products getProduct
// Return a single product from the data store
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrorProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

//Old method for listing all products
///func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
///	p.l.Println("Handle GET Products")
///
///	//fetch products from data storage
///	lp := data.GetProducts() //List of products
///	err := lp.ToJSON(rw)     //ToJSON is a method on the Products data object
///	if err != nil {
///		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
///		return
///	}
///}
