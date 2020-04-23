package newdata

import (
	"fmt"
	"time"
)

type Product struct {
	// the id for the coffee
	//
	// required: true
	// min: 1
	ID int `json:"id"`

	// the name for the coffee
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for the coffee
	//
	// required: false
	// max length: 10000
	Desc string `json:"desc"`

	// the price for the coffee
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"gt=0"`

	// the SKU for the coffee
	//
	// required: true
	// pattern: cof-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`

	// entry creation date
	//
	// required: false
	CreatedOn string `json:"-"`
	// entry last update date
	//
	// required: false
	UpdatedOn string `json:"-"`
	// entry deletion date
	//
	// required: false
	DeletedOn string `json:"-"`
}

type Products []*Product

var ErrorProductNotFound = fmt.Errorf("Product Not Found")

//returns the list of all the products from the mock data store
func GetProducts() Products {
	return productList
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func GetProductByID(id int) (*Product, error) {
	idx, err := findIndexByProductID(id)
	if err != nil {
		return nil, err
	}

	return productList[idx], nil
}

func AddProduct(p Product) {
	// get the next id in sequence
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, &p)
}

func UpdateProduct(p Product) error {
	idx, err := findIndexByProductID(p.ID)
	if err != nil {
		return err
	}

	// update the product in the DB
	productList[idx] = &p
	return nil
}

func DeleteProduct(id int) error {
	idx, err := findIndexByProductID(id)
	if err != nil {
		return err
	}

	//No idea how this works but it does. I blame slicetricks
	productList = append(productList[:idx], productList[idx+1:]...)
	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}

//finds the index of a product in the database
//returns -1 and an error if no product is found
func findIndexByProductID(id int) (int, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, nil
		}
	}
	return -1, ErrorProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
	&Product{
		ID:        1,
		Name:      "Espresso",
		Desc:      "Strong no milk",
		Price:     1.4,
		SKU:       "cof-espresso",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID:        2,
		Name:      "Latte",
		Desc:      "Frothy and milky",
		Price:     1.6,
		SKU:       "cof-latte",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID:        3,
		Name:      "Nescafe",
		Desc:      "Standard stuff",
		Price:     1.2,
		SKU:       "cof-nescafe",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
