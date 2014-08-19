package dnsraw

import (
	"encoding/binary"
	"math/rand"
	"time"
)

type Query struct {
	Id       uint16
	Flags    uint16
	Qdcount  uint16
	Ancount  uint16
	Nscount  uint16
	Arcount  uint16
	Arrcount uint16
	Request  *QueryRequest
}


// Creates a new Query
func NewQuery() *Query {
	return &Query{Id: genID(),
		Flags:    0x0100,
		Qdcount:  1,
		Ancount:  0,
		Nscount:  0,
		Arcount:  0,
		Arrcount: 0,
		Request:  NewQueryRequest()}
}

func genID() uint16 {
	rand.Seed(time.Now().UnixNano() * time.Now().UnixNano())
	return uint16(rand.Intn(65536-1) + 1)
}

func (q *Query) GenId() {
	q.Id = genID()
}

// SetId Sets the ID of the request
func (q *Query) SetId(t uint16) {
	q.Id = t
}

// SetRequest sets the type and name of the request
func (q *Query) SetRequest(s string, t string) {
	q.Request = NewQueryRequest()
	q.Request.SetClassDefault()
	q.Request.SetType(t)
	q.Request.SetName(s)
}

// Marshal buids the query into a byte array. The byte array is what will be sent
// as part of the DNS packet to the server.
func (q *Query) Marshal() []byte {
	//return byte array of query
	queryRequestb := q.Request.Marshal()
	idb := make([]byte, 2)
	flagsb := make([]byte, 2)
	qdcountb := make([]byte, 2)
	ancountb := make([]byte, 2)
	nscountb := make([]byte, 2)
	arcountb := make([]byte, 2)
	arrcountb := make([]byte, 2)
	binary.BigEndian.PutUint16(idb, q.Id)
	binary.BigEndian.PutUint16(flagsb, q.Flags)
	binary.BigEndian.PutUint16(qdcountb, q.Qdcount)
	binary.BigEndian.PutUint16(ancountb, q.Ancount)
	binary.BigEndian.PutUint16(nscountb, q.Nscount)
	binary.BigEndian.PutUint16(arcountb, q.Arcount)
	binary.BigEndian.PutUint16(arrcountb, q.Arrcount)

	b := make([]byte, 0)
	b = append(b, idb...)
	b = append(b, flagsb...)
	b = append(b, qdcountb...)
	b = append(b, nscountb...)
	b = append(b, arcountb...)
	b = append(b, arrcountb...)
	b = append(b, queryRequestb...)
	return b
}
