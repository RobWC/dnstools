package dnsraw

import (
	"encoding/binary"
	"strings"
)

//Default values for the various record types
var (
	ARecordType     = 0x0001
	NSRecordType    = 0x0002
	CNAMERecordType = 0x0005
	SOARecordType   = 0x0006
	WKSRecordType   = 0x000B
	PTRRecordType   = 0x000C
	MXRecordType    = 0x000F
	SRVRecordType   = 0x0021
	A6RecordType    = 0x0026
	ANYRecordType   = 0x00FF
)

//Default value for the query class
// This is the default and it is not changed, unless for fun purposes
var (
	INClass = 0x0001
)

//DNSQuestion struct sets up the elments needed to generate a DNS Question
type DNSQuestion struct {
	Name  string
	Type  uint16
	Class uint16
}

//NewDNSQuestion returns an empty and initalized DNSQuestion
func NewDNSQuestion() *DNSQuestion {
	return &DNSQuestion{Name: "",
		Type:  0,
		Class: 0}
}

//SetType specifies the typq of the query
// Accepts the following
// a for A Record
// ns for NS Record
// cname for CNAME Record
// mx for MX Record
// soa for SOA Record
// wks for WKS Record
// ptr for PTR Record
// srv for SRV Record
// a6 for A6 Record
// any for ANY Record
func (qr *DNSQuestion) SetType(s string) {
	s = strings.ToLower(s)
	switch {
	case s == "a":
		qr.Type = ARecordType
	case s == "ns":
		qr.Type = NSRecordType
	case s == "cname":
		qr.Type = CNAMERecordType
	case s == "mx":
		qr.Type = MXRecordType
	case s == "soa":
		qr.Type = SOARecordType
	case s == "wks":
		qr.Type = WKSRecordType
	case s == "ptr":
		qr.Type = PTRRecordType
	case s == "srv":
		qr.Type = SRVRecordType
	case s == "a6":
		qr.Type = A6RecordType
	case s == "any":
		qr.Type = ANYRecordType
	}
}

//SetName to be encoded into the query
func (qr *DNSQuestion) SetName(s string) {
	qr.Name = s
}

//SetClass specify a differnt class name
func (qr *DNSQuestion) SetClass(i uint16) {
	qr.Class = i
}

//SetClassDefault sets the DNS question class to the default value
func (qr *DNSQuestion) SetClassDefault() {
	qr.Class = INClass
}

//Marshal returns the DNSQuestion in a network ready binary slice
func (qr *DNSQuestion) Marshal() []byte {
	//return byte array

	//break apart the name by period
	//count eatch segment
	var nameBytes []byte

	splitName := strings.Split(qr.Name, ".")
	for name := range splitName {
		nameBytes = append(nameBytes, byte(len(splitName[name])))
		nameBytes = append(nameBytes, []byte(splitName[name])...)
	}

	nameBytes = append(nameBytes, 0)
	//nameLen := len(nameBytes)
	var b []byte
	typeb := make([]byte, 2)
	classb := make([]byte, 2)
	binary.BigEndian.PutUint16(typeb, qr.Type)
	binary.BigEndian.PutUint16(classb, qr.Class)
	b = append(b, nameBytes...)
	b = append(b, typeb...)
	b = append(b, classb...)
	return b
}
