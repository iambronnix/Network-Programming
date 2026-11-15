package mime

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestMultipartPost(t *testing.T){
	reqBody := new(bytes.Buffer)
	w := multipart.NewWriter(reqBody)

	for k,v := range map[string]string{
		"date" : time.Now().Format(time.RFC3339),
		"description": "form values with attached files",
	}{
		err := w.WriteField(k, v)
		if err != nil{
			log.Fatal(err)
		}
	}
	for i, file := range []string{
		"./files/hello.txt",
		"./files/goodbye.txt",
	}{
		filePart, err := w.CreateFormFile(fmt.Sprintf("file%d", i+1),
        filepath.Base(file))
		if err != nil{
			t.Fatal(err)
		}
		f, err := os.Open(file)
		if err != nil{
			t.Fatal(err)
		}

		_, err = io.Copy(filePart,f)
		_ = f.Close()
		if err != nil{
			t.Fatal(err)
		}
	}
	err := w.Close()
	if err != nil{
		t.Fatal(err)
	}

}