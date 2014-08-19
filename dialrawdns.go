package dialrawdns

/*
con, err := net.ListenPacket("ip4:udp", "0.0.0.0")
if err != nil {
    log.Fatalln(err)
}

//new raw packet connection
rawCon, err := ipv4.NewRawConn(con)
if err != nil {
    log.Fatalln(err)
}

//set final payload

//set packet length
rdns.IPHeaders.TotalLen = 20 + len(queryb) + len(udpHead)
rdns.Payload = make([]byte, 0)
rdns.Payload = append(rdns.Payload, udpHead...)
rdns.Payload = append(rdns.Payload, queryb...)

rawCon.WriteTo(rdns.IPHeaders, rdns.Payload, rdns.CtrlMsg)
*/
