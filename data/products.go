package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"io"
	"regexp"
	"time"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name" validate:"required"`
	Desc      string  `json:"desc"`
	Price     float32 `json:"price" validate:"gt-0"`
	SKU       string  `json:"sku" validate:"required,sku"`
	CreatedOn string  `json:"-"`
	UpdatedOn string  `json:"-"`
	DeletedOn string  `json:"-"`
}

type Products []*Product

var ErrorProductNotFound = fmt.Errorf("Product Not Found")

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(`cof-[a-z]+`)
	matches := reg.FindAllString(fl.Field().String(), -1)
	if len(matches) == 1 {
		return false
	}
	return true
}

func (p *Products) ToJSON(w io.Writer) error {
	err := json.NewEncoder(w)
	return err.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	err := json.NewDecoder(r)
	return err.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p
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
