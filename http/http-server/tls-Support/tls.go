package tlssupport

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main(){
	srv := &http.Server{
		Addr: "127.0.0.1:8443",
		Handler: mux,
		IdleTimeout: 5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}
	l, err := net.Listen("tcp", srv.Addr)
	if err != nil{
		log.Fatal(err)
	}

	go func(){
		err := srv.ServeTLS(l, "cert.pem", "key.pem")//for production use https://letsencrypt.org/
		                                             //for development check out https://github.com/FiloSottile/mkcert/
		if err != http.ErrServerClosed{
			log.Fatal(err)
			
		}
	}()
}