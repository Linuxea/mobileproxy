package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {

	port := flag.String("port", "8888", "port")
	flag.Parse()
	fmt.Println("port is", *port)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":"+*port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})))

}
