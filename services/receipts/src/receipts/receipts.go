package main

import (
	"encoding/json"
	"net/http"
	"routes"

	"github.com/go-pg/pg"
)

func createName(db *pg.DB) routes.HandlerFunc {
	type Request struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	}

	type Payload struct {
		Status string `json:"status"`
		Data   Name   `json:"data"`
	}
	return func(w http.ResponseWriter, r *routes.Request) {
		req := &Request{}

		json.NewDecoder(r.Body).Decode(&req)

		name := &Name{
			Name: req.Name,
			Key:  req.Key,
		}
		err := db.Insert(name)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(name)
	}
}
