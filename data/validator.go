package data

import (
	"github.com/go-playground/validator"
	"regexp"
)

//Here are a couple methods I prefer using instead of their
//other versions (added a WRONG to those invalidation.go)

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(`cof-[a-z]+`)
	matches := reg.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}
