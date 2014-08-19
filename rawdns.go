package dnstools

import (
	"net"
	"strconv"
	"strings"

	"code.google.com/p/go.net/ipv4"
)

//RawDNS data struct. Contains L3, L4 Headers, payload and control message for specifying the egress interface
// BUG(robwc) Only supports IPv4 today
type RawDNS struct {
	IPHeaders     *ipv4.Header
	UDPHeader     *UDPHeader
	DNSQuestion   *DNSQuestion
	LocalAddress  net.IP
	RemoteAddress net.IP
	Payload       []byte
	CtrlMsg       *ipv4.ControlMessage
}

//NewRawDNS return an initalized RawDNS struct
func NewRawDNS() *RawDNS {
	return &RawDNS{IPHeaders: new(ipv4.Header),
		UDPHeader:     new(UDPHeader),
		LocalAddress:  net.IPv4(0, 0, 0, 0),
		RemoteAddress: net.IPv4(0, 0, 0, 0),
		Payload:       make([]byte, 0),
		CtrlMsg:       new(ipv4.ControlMessage)}
}

//SetDestPort set destination port
func (rdns *RawDNS) SetDestPort(port uint16) {
	rdns.UDPHeader.SetDstPort(port)
}

//SetLocalAddress set local or source address
func (rdns *RawDNS) SetLocalAddress(ip string) {
	parsedIP := strings.Split(ip, ".")
	ip0, _ := strconv.Atoi(parsedIP[0])
	ip1, _ := strconv.Atoi(parsedIP[1])
	ip2, _ := strconv.Atoi(parsedIP[2])
	ip3, _ := strconv.Atoi(parsedIP[3])
	rdns.LocalAddress = net.IPv4(byte(ip0), byte(ip1), byte(ip2), byte(ip3))
}

//SetRemoteAddress set remote address of
func (rdns *RawDNS) SetRemoteAddress(ip string) {
	parsedIP := strings.Split(ip, ".")
	ip0, _ := strconv.Atoi(parsedIP[0])
	ip1, _ := strconv.Atoi(parsedIP[1])
	ip2, _ := strconv.Atoi(parsedIP[2])
	ip3, _ := strconv.Atoi(parsedIP[3])
	rdns.RemoteAddress = net.IPv4(byte(ip0), byte(ip1), byte(ip2), byte(ip3))
	rdns.CtrlMsg.Dst = rdns.RemoteAddress
}

//SetUDPHeader sets the UDP header for the packet
func (rdns *RawDNS) SetUDPHeader(header UDPHeader) {
	//set the UDP headers
	rdns.UDPHeader = header
	rdns.UDPHeader.GenRandomSrcPort()
	rdns.UDPHeader.SetChecksum(0)
}

//SetDNSQuestion set the payload of the DNS packet
func (rdns *RawDNS) SetDNSQuestion(question DNSQuestion) {
	rdns.DNSQuestion = question
}

//Marshall return the items required to send a raw packet
// returns the three elements needed to sent into a raw packet
// IPheaders []byte, Payload []byte, ControlMessage ipv4.ControlMessage
func (rdns *RawDNS) Marshall() ([]byte, []byte, ipv4.ControlMessage) {

	//set the IP headers
	rdns.IPHeaders.Src = rdns.LocalAddress
	rdns.IPHeaders.Dst = rdns.RemoteAddress
	rdns.IPHeaders.Protocol = IPProtoUDP
	rdns.IPHeaders.Len = IPHeaderLen
	rdns.IPHeaders.Version = 4
	rdns.IPHeaders.TTL = 128

	//set the UDP header length
	rdns.UDPHeader.SetLen(8 + uint16(len(queryb)))
	udpHead, _ := rdns.UDPHeader.Marshal()

	//set the query
	queryb := rdns.DNSQuestion.Marshal()

	//set the control message
	rdns.CtrlMsg.TTL = 128
	rdns.CtrlMsg.IfIndex = config.Interface.Index

	//set final payload

	//set packet length
	rdns.IPHeaders.TotalLen = 20 + len(queryb) + len(udpHead)
	rdns.Payload = make([]byte, 0)
	rdns.Payload = append(rdns.Payload, udpHead...)
	rdns.Payload = append(rdns.Payload, queryb...)

	return rdns.IPHeaders, rdns.Payload, rdns.CtrlMsg

}
