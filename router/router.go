package router

import (
	"context"
	"encoding/json"
	"fmt"
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


	data,err := d.GetAll(client, "zips")

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

	js ,_ := json.Marshal(data)

	fmt.Fprint(w, string(js))
}

func ApiGet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	q, err := url.Parse(r.URL.String())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if q.RawQuery == "" {
		http.Error(w, "No query string", http.StatusBadRequest)
		return
	}

	query := q.Query()

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

	client,err := d.NewClient()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer client.Disconnect(context.TODO())


	ret := []s.Zip{}

	for _,v := range query["value"] {
		data,err := d.GetObjs(client, "zips", query["type"][0], v )
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ret = append(ret, data...)
	}

	log.Println(query, len(query["value"]))


	data, err := json.Marshal(ret)

	fmt.Fprint(w,string(data))
}

func ApiPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := r.GetBody()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	log.Println(data)

}
