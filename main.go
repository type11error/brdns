package main

import (
	"log"
	"github.com/miekg/dns"
)

// handleDNSRequest processes incoming DNS queries
func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	for _, q := range r.Question {
		switch q.Qtype {
		case dns.TypeA: // Handle A record queries
			if q.Name == "example.com." {
				rr, err := dns.NewRR("example.com. 3600 IN A 127.0.0.1")
				if err == nil {
					msg.Answer = append(msg.Answer, rr)
				}
			}
		case dns.TypeMX: // Handle MX record queries
			if q.Name == "example.com." {
				rr, err := dns.NewRR("example.com. 3600 IN MX 10 mail.example.com.")
				if err == nil {
					msg.Answer = append(msg.Answer, rr)
				}
			}
		default:
			log.Printf("Unhandled query type: %d", q.Qtype)
		}
	}

	if err := w.WriteMsg(&msg); err != nil {
		log.Printf("Failed to write message: %v", err)
	}
}

func main() {
	dns.HandleFunc("example.com.", handleDNSRequest)

	server := &dns.Server{
		Addr: ":8053", // Listen on port 8053
		Net:  "udp",   // Use UDP
	}

	log.Printf("Starting DNS server on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

