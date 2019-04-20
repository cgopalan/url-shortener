package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

const prefix string = "cgo.pl"

const store = ""

var currentStore urlStore

func generate(url string) string {
	if val, ok := currentStore.lookUpLongURL(url); ok {
		return fmt.Sprintf("%v/%v", prefix, val)
	}
	su := generateShortURL()
	currentStore.addToStore(su, url)
	return fmt.Sprintf("%v/%v", prefix, su)
}

func generateShortURL() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
	const lenLetters = len(letterBytes)
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(lenLetters)]
	}
	return string(b)
}

func lookUpURL(surl string) string {
	return currentStore.lookUp(surl)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		keys, _ := r.URL.Query()["url"]
		fmt.Fprintf(w, generate(keys[0]))
	})

	http.HandleFunc("/unmask", func(w http.ResponseWriter, r *http.Request) {
		keys, _ := r.URL.Query()["url"]
		fmt.Fprintf(w, lookUpURL(keys[0]))
	})

	currentStore = memStore{}
	if store == "REDIS" {
		currentStore = redisStore{}
		fmt.Println("Current store is REDIS")
	} else {
		fmt.Println("Current store is memory")
	}
	currentStore.init()

	fmt.Println("Starting server.")
	http.ListenAndServe(":8010", nil)
}
