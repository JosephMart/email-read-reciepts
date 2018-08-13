package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-pg/pg"
)

func createName(db *pg.DB) http.Handler {
	type Request struct {
		Name string `json:"name"`
		Key  string `json:"key"`
	}

	type Payload struct {
		Status string `json:"status"`
		Data   Name   `json:"data"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
}
