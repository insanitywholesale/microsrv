package main

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"log"
	"microsrv/data"
	"microsrv/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "products-api: ", log.LstdFlags) //logger to be used for the products api
	v := data.NewValidation()
	redocOpts := middleware.RedocOpts{SpecURL: "/swagger.yml"} //Redoc option to use YAML instead of JSON

	ph := handlers.NewProducts(l, v)       //Product handler
	sh := middleware.Redoc(redocOpts, nil) //Swagger handler

	//responseweriter and request get passed automatically to the function in HandleFunc
	sm := mux.NewRouter()                                        //Gorilla ServeMux aka root Router
	getRouter := sm.Methods(http.MethodGet).Subrouter()          //Add a SubRouter to the root Router
	getRouter.HandleFunc("/products", ph.GetProducts)            //Assign function to path
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.GetProduct) //Assing function to path

	getRouter.Handle("/docs", sh)                                     //Handle the /docs path on GET requests
	getRouter.Handle("/swagger.yml", http.FileServer(http.Dir("./"))) //Serve swagger.yml so the above works

	postRouter := sm.Methods(http.MethodPost).Subrouter() //Add another Subrouter to the root Router
	postRouter.HandleFunc("/products", ph.AddProduct)     //Assign function to endpoint
	postRouter.Use(ph.MiddlewareValidateProduct)          //Middleware to validate json structure

	putRouter := sm.Methods(http.MethodPut).Subrouter() //Add another Subrouter to the root Router
	putRouter.HandleFunc("/products", ph.UpdateProduct) //Assign function to path
	putRouter.Use(ph.MiddlewareValidateProduct)         //Middleware to validate json structure

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()          //Add another Subrouter to the root Router
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct) //Assign function to path

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  100 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Println("Starting server")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	l.Println("Received terminate, starting graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 11*time.Second) //Timeout context that waits up to 8 seconds
	s.Shutdown(tc)
}
