package dnstools

import (
	//"lib/channels/chanman"
	"log"
	"net"
	"sync"

	"./dnsstat"
	"github.com/miekg/dns"
)

func DnsQuery(query string, wg *sync.WaitGroup, config *dns.ClientConfig, cm *chanman.ChanMan) {
	defer wg.Done()
	dnsClient := new(dns.Client)
	message := new(dns.Msg)
	message.SetQuestion(dns.Fqdn(query), dns.TypeA)
	message.RecursionDesired = true

	response, rtt, err := dnsClient.Exchange(message, net.JoinHostPort(config.Servers[0], "53"))

	if err != nil {
		log.Println(err)
		cm.RunChan <- true
	}

	if response == nil {
	} else {
		if response.Rcode != dns.RcodeSuccess {
			log.Println("query fail")
		}
		var stat = new(dnsstat.Info)
		stat.Rtt = rtt
		cm.StatsChan <- stat
		cm.RunChan <- true
	}
}
