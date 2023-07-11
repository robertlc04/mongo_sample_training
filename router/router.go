package router

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/julienschmidt/httprouter"

	d "github.com/mongo_sample_training/database"
	s "github.com/mongo_sample_training/structs"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "public/index.html")
}

func PostPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "public/add.html")
}

func ApiGetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	x, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(x.RawQuery)

	client, err := d.NewClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer client.Disconnect(context.TODO())

	data, err := d.GetAll(client, "zips")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func ApiGet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	q, err := url.Parse(r.URL.String())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Empty query error
	if q.RawQuery == "" {
		http.Error(w, "No query string", http.StatusBadRequest)
		return
	}

	query := q.Query()

	// Errors Control
	if _, err := query["type"]; !err {
		http.Error(w, "Please Specify the type\n Example: /api/search?type=city&value=Rome", http.StatusInternalServerError)
		return
	}

	if len(query["type"]) != 1 {
		http.Error(w, "Please Specify one type\n Example: /api/search?type=city&value=Rome", http.StatusInternalServerError)
		return
	}

	if _, err := query["value"]; !err {
		http.Error(w, "Please Specify the value to the query\n Example: /api/search?type=city&value=Rome", http.StatusInternalServerError)
		return
	}

	if len(strings.Join(query["value"], " ")) < 1 {
		http.Error(w, "Please Specify the value to the query\n Example: /api/search?type=city&value=Rome", http.StatusInternalServerError)
		return
	}

	// Database Connect
	client, err := d.NewClient()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer client.Disconnect(context.TODO())

	ret := []s.Zip{}

	for _, v := range query["value"] {
		data, err := d.GetObjs(client, "zips", query["type"][0], v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ret = append(ret, data...)
	}

	json.NewEncoder(w).Encode(ret)
}

func ApiPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	var body map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(body)

	w.Header().Set("Content-Type", "application/json")

	resp := map[string]string{ "status": "ok"}

	json.NewEncoder(w).Encode(resp)

}
