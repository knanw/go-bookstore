package main

import (
	"fms"
	"log"
	"net/http"

	"github.com/akhil/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/jinzu/gorm/dialects/mysql"
	"github.com/knanw/go-bookstore/pkg/routes"
)

func main()  {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) //create server on port 9010
}