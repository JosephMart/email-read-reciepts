package main

import (
	"fmt"
	"image/png"
	"log"
	"net"
	"net/http"
	"os"
	"routes"

	"github.com/go-pg/pg"
)

func getIP(req *routes.Request) string {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)

	if err != nil {
		fmt.Println("userip: %q is not IP:port", req.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		//return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
		// fmt.Println("userip: %q is not IP:port", req.RemoteAddr)
		return ""
	}

	// This will only be defined when site is accessed via non-anonymous proxy
	// and takes precedence over RemoteAddr
	// Header.Get is case-insensitive
	forward := req.Header.Get("X-Forwarded-For")

	fmt.Println("IP: ", ip)
	fmt.Println("Forwarded for:", forward)

	return ip
}

func serveImg(db *pg.DB) routes.HandlerFunc {
	return func(w http.ResponseWriter, r *routes.Request) {
		key := r.Kwargs["key"]

		name := new(Name)
		err := db.Model(name).Where("key = ?", key).Select()

		ip := getIP(r)

		if err != nil {
			fmt.Println(err)
		} else {
			event := &Event{Ip: ip}
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
