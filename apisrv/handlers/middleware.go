package handlers

import (
	"context"
	"microsrv/apisrv/data"
	"net/http"
)

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		//err := data.FromJSON(prod, r.Body) //suggested method but doesn't work
		if err != nil {
			p.l.Println("[ERROR]: deserializing product", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		err = prod.Validate()
		//if errs != nil {
		//	p.l.Println("[ERROR]: validating product", err)

		//	rw.WriteHeader(http.StatusUnprocessableEntity)
		//	data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
		//	return
		//}

		//add the procuct to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		//call the next handler(could be more middleware or final handler)
		next.ServeHTTP(rw, req)
	})
}
