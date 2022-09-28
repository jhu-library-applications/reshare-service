package main

import (
	"io"
	"log"
	"net/http"
	"reshare-service/reshare"
)

func main() {
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		var request reshare.Request

		isn, isnPresent := r.URL.Query()["isn"]
		title, titlePresent := r.URL.Query()["title"]
		author, authorPresent := r.URL.Query()["author"]

		// Set the ISN
		if !isnPresent || len(isn[0]) < 1 {
			log.Println("URL Param 'isn' is missing")

			// If the ISN is missing set the title. If the title is missing then return.
			// Title is the minimum to make a request
			if !titlePresent || len(title[0]) < 1 {
				log.Println("URL Param 'title' is missing")
				return
			} else {
				request.Title = title[0]

				if authorPresent {
					request.Author = author[0]
				}
			}
		} else {
			request.Isn = isn[0]

			if titlePresent {
				request.Title = title[0]
			}

			if authorPresent {
				request.Author = author[0]
			}
		}

		io.WriteString(w, request.ItemRequest())
	})

	http.ListenAndServe(":5050", nil)
}
