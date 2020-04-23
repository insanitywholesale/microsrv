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

	params := products.NewGetProductsParams()
	prod, err := c.Products.GetProducts(params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(prod)
	fmt.Printf("%#v", prod.GetPayload()[0])
	//t.Fail()
}
