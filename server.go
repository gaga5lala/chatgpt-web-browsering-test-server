package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Log the request
	fmt.Println("Headers:", r.Header)
	fmt.Println("URL:", r.URL.Path)
	fmt.Println("Query:", r.URL.RawQuery)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}
	fmt.Println("Body:", string(body))

	// Respond with status code 200 and a JSON payload
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "This is a JSON response"}`))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "robots.txt")
	})
	// server sitemap.xml
	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "sitemap.xml")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	// go func() {
	// log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
	// }()
}
