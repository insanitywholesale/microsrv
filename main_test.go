package main

import (
	"fmt"
	"microsrv/client/client"
	"microsrv/client/client/products"
	"testing"
)

func TestClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	prodParams := products.NewGetProductParams()
	prod, err := c.Products.GetProduct(prodParams)
	if err != nil {
		t.Fatal(err)
	}

	prodsParams := products.NewGetProductsParams()
	prods, errs := c.Products.GetProducts(prodsParams)
	if errs != nil {
		t.Fatal(err)
	}

	fmt.Println(prod)
	fmt.Println(prods)
	fmt.Printf("%#v", prod.GetPayload())
	fmt.Println("")
	fmt.Printf("%#v", prods.GetPayload()[3])
	fmt.Println("")

	//Remove if you don't want to debug
	//Aka if you actually want to test
	//Lmao
	//Nobody tests
	t.Fail()
}
