package main

import "fmt"
import "log"
import "github.com/google/tcpproxy"

func main() {
	var p tcpproxy.Proxy
	p.AddRoute(":5000", tcpproxy.To("localhost:16000"))
	p.AddRoute(":5001", tcpproxy.To("localhost:16001"))
	fmt.Println("Ready to accept requests")
	log.Fatal(p.Run())
}




