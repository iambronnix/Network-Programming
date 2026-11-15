package main

import (
	"log"
	"net/http"
	"time"
)

func headTime(){
	resp, err := http.Head("http://www.time.gov/")
	if err != nil{
		log.Fatal(err)
	}
	defer func(){
	_ = resp.Body.Close()
	}()

	now := time.Now().Round(time.Second)
	date := resp.Header.Get("Date")
	if date == ""{
		log.Fatal("no Date header received from time.gov")
	}

	dt, err := time.Parse(time.RFC1123, date)
	if err != nil{
		log.Fatal(err)
	}
	log.Printf("time.gov: %s (skew %s)", dt, now.Sub(dt))
}