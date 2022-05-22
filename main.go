package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	_ "embed"

	"github.com/elazarl/goproxy"
	"github.com/magisterquis/connectproxy"
	"golang.org/x/net/proxy"
)

//go:embed white
var whites string

func main() {

	do := doProxy()
	no := noProxy()

	whiteList := strings.Split(whites, ",")
	for _, v := range whiteList {
		fmt.Println(v)
	}

	log.Fatal(http.ListenAndServe(":8888", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for index := range whiteList {
			if whiteList[index] != "" && strings.Contains(r.Host, whiteList[index]) {
				fmt.Println(r.Host + "不走代理")
				no.ServeHTTP(w, r)
				return
			}
		}

		fmt.Println(r.Host + "走代理")
		do.ServeHTTP(w, r)

	})))

}

func doProxy() http.Handler {

	hs := goproxy.NewProxyHttpServer()

	u, _ := url.Parse("http://127.0.0.1:50475")
	d, _ := connectproxy.New(u, proxy.Direct)
	hs.Tr = &http.Transport{
		Dial: d.Dial,
	}
	hs.Verbose = true

	return hs
}

func noProxy() http.Handler {

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	return proxy
}
