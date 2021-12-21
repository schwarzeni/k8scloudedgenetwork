package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/miekg/dns"
)

//var records = map[string]string{
//	"test.edge.zone.": "10.211.55.2",
//}
//
//func parseQuery(m *dns.Msg) {
//	for _, q := range m.Question {
//		switch q.Qtype {
//		//case dns.TypeA:
//		default:
//			log.Printf("Query for %s\n", q.Name)
//			ip := records[q.Name]
//			if ip != "" {
//				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
//				if err == nil {
//					m.Answer = append(m.Answer, rr)
//				}
//			}
//		}
//	}
//}

var cloudEPIP string = "10.211.55.2"

func parseQuery(m *dns.Msg) {
	for _, q := range m.Question {
		switch q.Qtype {
		//case dns.TypeA:
		default:
			// TODO: 这里默认请求格式都是正确的
			log.Printf("Query for %s\n", q.Name)
			rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, cloudEPIP))
			if err == nil {
				m.Answer = append(m.Answer, rr)
			}
		}
	}
}

func handleDnsRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m)
	}

	fmt.Println(m.Answer)
	w.WriteMsg(m)
}

func main() {
	if len(os.Args) > 1 {
		cloudEPIP = os.Args[1]
	}
	// attach request handler func
	dns.HandleFunc(".", handleDnsRequest)

	// start server
	port := 53530
	server := &dns.Server{Addr: ":" + strconv.Itoa(port), Net: "udp"}
	log.Printf("Starting at %d\n", port)
	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n ", err.Error())
	}
}
