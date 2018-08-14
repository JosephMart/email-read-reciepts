package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"routes"

	"github.com/go-pg/pg"
)

func serveImg(db *pg.DB) routes.HandlerFunc {
	return func(w http.ResponseWriter, r *routes.Request) {
		key := r.Kwargs["key"]

		name := new(Name)
		err := db.Model(name).Where("key = ?", key).Select()

		if err != nil {
			fmt.Println(err)
		} else {
			event := &Event{}
			err := db.Insert(event)

			if err != nil {
				fmt.Println(err)
			}

			name.Events = append(name.Events, event)
			err = db.Update(name)

			if err != nil {
				fmt.Println(err)
			}
		}

		Path := "./img.png"

		file, err := os.Open(Path)
		if err != nil {
			log.Fatal(err)
		}

		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		png.Encode(w, img)
	}
}
