package newhandlers

import (
	data "microsrv/newdata"
	"net/http"
)

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	//fetch products from data storage
	lp := data.GetProducts() //List of products
	err := lp.ToJSON(rw)     //ToJSON is a method on the Products data object
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
