package main

type urlStore interface {
	init()
	lookUp(surl string) string
	addToStore(url string, su string)
	lookUpLongURL(url string) (string, bool)
}
