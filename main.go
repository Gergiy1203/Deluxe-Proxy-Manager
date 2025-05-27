package main

import (
     "cryto/tls"
	  "fmt"
	  "io"
	  "log"
	  "net"
	  "net/http"
)

const proxyListenAdddress = "0.0.0.0:3128"

func main() {
	Addr:       proxyListenAdddress,
	Handler:    http.HandlerFunc(connectHadler),
	TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){},
}
log.Println("Started proxy at:", proxyServer.Addr)
if err := proxyServer.ListenAndServer(); err != nil {
	log.Println("Server failed:", err)
}
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method 
}