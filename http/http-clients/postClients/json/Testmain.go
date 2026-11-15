package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)
type User struct{
	First, 
	Last string
}

func handlePostUser(t *testing.T)func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		defer func(r io.ReadCloser){
			_, _ = io.Copy(io.Discard, r)
			_ = r.Close()
		}(r.Body)

		if r.Method != http.MethodPost{
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}
		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil{
			log.Print(err)
			http.Error(w, "decode failed", http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusAccepted)
	}
}

func postUser(t *testing.T){
	ts := httptest.NewServer(http.HandlerFunc(handlePostUser(t)))
	defer ts.Close()
	 resp, err := http.Get(ts.URL)
		if err != nil{
			log.Fatal(err)
		}
		if resp.StatusCode != http.StatusMethodNotAllowed{
			log.Fatalf("expected status %d; actual status %d", http.StatusMethodNotAllowed, resp.StatusCode)
		}

		buf := new(bytes.Buffer)
		u := User{First: "erick", Last: "Mwendwa"}
		err = json.NewEncoder(buf).Encode(&u)
		if err != nil{
			log.Fatal(err)
		}

		resp, err = http.Post(ts.URL, "application/json", buf)
		if err != nil{
			log.Fatal(err)
		}

		if resp.StatusCode != http.StatusAccepted {
			log.Fatalf("expected status %d; actual status %d ", http.StatusAccepted,resp.StatusCode)
		}
		_ = resp.Body.Close()
}