package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	r "github.com/mongo_sample_training/router"
)

func main() {
	
	router := httprouter.New()

	// Pages
	router.GET("/", r.Index)
	router.GET("/api/add", r.PostPage)
	
	// Json Data
	router.GET("/api/all", r.ApiGetAll)
	router.GET("/api/search", r.ApiGet)
	router.POST("/api/add", r.ApiPost)

	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
