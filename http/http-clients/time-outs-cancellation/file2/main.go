package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func blockIndefinitely(w http.ResponseWriter, r *http.Request){
	select{}//blocking effect
}
func blockIndefinitelyWithTimeout(){
	ts := httptest.NewServer(http.HandlerFunc(blockIndefinitely))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,ts.URL, nil)
	if err != nil{
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		if !errors.Is(err, context.DeadlineExceeded){
			log.Fatal(err)
		}
		return
	}
	_ = resp.Body.Close()
}