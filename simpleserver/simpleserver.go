package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

type ProxyHandler struct{}

func (p *ProxyHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.String())
	host, port, err := net.SplitHostPort(req.Host)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(host, port)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", &ProxyHandler{}))
}
