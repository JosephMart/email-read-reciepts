package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"
)

func serveImg() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Path := "./img.png"
		var ImageExtension = strings.Split(Path, ".png")
		var ImageNum = strings.Split(ImageExtension[0], "/")
		var ImageName = ImageNum[len(ImageNum)-1]
		fmt.Println(ImageName)
		file, err := os.Open(Path)
		if err != nil {
			log.Fatal(err)
		}

		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		png.Encode(w, img) // Write to the ResponseWriter
	})
}
