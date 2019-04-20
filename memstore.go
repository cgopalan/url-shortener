package main

import "fmt"

var urlMap map[string]string
var longURLMap map[string]string

type memStore struct {
}

func (m memStore) init() {
	urlMap = make(map[string]string)
	longURLMap = make(map[string]string)
}

func (m memStore) addToStore(su string, url string) {
	urlMap[su] = url
	longURLMap[url] = su
}

func (m memStore) lookUp(surl string) string {
	if val, ok := urlMap[surl[7:]]; ok {
		return val
	}
	return fmt.Sprintf("%v cannot be resolved", surl)
}

func (m memStore) lookUpLongURL(url string) (string, bool) {
	val, ok := longURLMap[url]
	return val, ok
}
