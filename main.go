package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "time"	
//	"github.com/gin-gonic/gin"
//	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)


func main() {
	log.Printf("started")
//	uri := fmt.Sprintf("mongodb://tal:tal@104.197.29.173:27017/test&authorization=disabled")
//	session, err := mgo.Dial(uri)
//	log.Printf("dial")
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()
//	log.Printf("close")
	
	port := "8443"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	tlsCert, tlsKey := os.Getenv("TLS_CERT"), os.Getenv("TLS_KEY")
	if tlsCert == "" {
		log.Fatal("TLS_CERT environment variable must be set")
	}
	if tlsKey == "" {
		log.Fatal("TLS_KEY environment variable must be set")
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	server.HandleFunc("/", hello)

	// start the web server on port and accept requests
	log.Printf("tls cert: %s", tlsCert)
	log.Printf("tls key: %s", tlsKey)
	log.Printf("Server listening on port %s", port)
	err2 := http.ListenAndServeTLS(":"+port, tlsCert, tlsKey, server)
	log.Fatal(err2)		
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hi, BioCatch!\n")
	fmt.Fprintf(w, "Protocol: %s!\n", r.Proto)
	fmt.Fprintf(w, "Hostname: %s\n", host)
	if headerIP := r.Header.Get("X-Forwarded-For"); headerIP != "" {
		fmt.Fprintf(w, "Client IP (X-Forwarded-For): %s\n", headerIP)
	}
}

