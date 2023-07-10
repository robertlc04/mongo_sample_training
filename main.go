package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	r "github.com/mongo_sample_training/router"
)

func main() {
	
	router := httprouter.New()

	router.GET("/", r.Index)
	router.GET("/api", r.ApiGetAll)
	router.GET("/api/search", r.ApiGet)
	router.POST("/api", r.ApiPost)

	fmt.Print("HOLA\n")

	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
