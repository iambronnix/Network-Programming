package main

import (
	"log"
	"net/http"
	"net/http/httptest"
)

func blockIndefinitely(w http.ResponseWriter, r *http.Request){
	select{}//blocking effect
}
func handleblock(){
	ts := httptest.NewServer(http.HandlerFunc(blockIndefinitely))
	_, _ = http.Get(ts.URL)
	log.Fatal("client did not indefinitely block")
}