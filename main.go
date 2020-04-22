package main

import (
	"context"
	"github.com/gorilla/mux"
	//"github.com/insanitywholesale/microsrv/handlers"
	"github.com/go-openapi/runtime/middleware"
	"log"
	handlers "microsrv/newhandlers" //offline version of github import
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "products-api: ", log.LstdFlags)   //logger to be used for the products api
	redocOpts := middleware.RedocOpts{SpecURL: "/swagger.yml"} //Redoc option to use YAML instead of JSON

	ph := handlers.NewProducts(l)          //Product handler
	sh := middleware.Redoc(redocOpts, nil) //Swagger handler

	sm := mux.NewRouter()                               //Gorilla ServeMux aka root Router
	getRouter := sm.Methods(http.MethodGet).Subrouter() //Add a SubRouter to the root Router
	getRouter.HandleFunc("/", ph.GetProducts)           //(responsewriter and request are passed automatically)

	getRouter.Handle("/docs", sh)                                     //Handle the /docs path on GET requests
	getRouter.Handle("/swagger.yml", http.FileServer(http.Dir("./"))) //Serve swagger.yml

	postRouter := sm.Methods(http.MethodPost).Subrouter() //Add another Subrouter to the root Router
	postRouter.HandleFunc("/", ph.AddProduct)             //(responsewriter and request are passed automatically)
	postRouter.Use(ph.MiddlewareValidateProduct)          //Middleware to validate json structure

	putRouter := sm.Methods(http.MethodPut).Subrouter()    //Add another Subrouter to the root Router
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct) //(responsewriter and request are passed automatically)
	putRouter.Use(ph.MiddlewareValidateProduct)            //Middleware to validate json structure

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter() //Add another Subrouter to the root Router
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct) //responsewriter and request are passed automatically
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