package main

import (
	"log"
	"net"
	"sip"
	"strings"
)

var s sip.SipClient
var appconfig AppConfiguration

func Register() {
	s = sip.CreateClient()
	s.SetDefaultTransport(appconfig.Transport)

	proxyIP := appconfig.ProxyIP
	localIP := GetOutboundIP(proxyIP)
	log.Println("localIP: ", localIP)
	r := sip.RegisterInfo{
		sip.Connectinfo{
			appconfig.Transport,
			proxyIP,
			appconfig.ProxyPort,
		},
		sip.Connectinfo{
			appconfig.Transport,
			localIP,
			appconfig.LocalPort,
		},
		appconfig.Phonenumber,
		sip.DigestUserInfo(appconfig.Username, appconfig.Password),
		0,
	}

	//---
	result, err := s.TryRegister(&r)

	log.Println("Register result:", result, "Error: ", err)

	s.Listen("tcp", "0.0.0.0", appconfig.LocalPort)
}

func main() {
	var err error
	forever := make(chan bool)
	appconfig, err = ReadConfig()
	if err != nil {
		log.Println("ERROR reading config file: ", err)
	}

	Register()
	Init(appconfig.PinNumber)
	s.OnIncomingCall(func(call *sip.Call) {
		log.Println("Incoming Call")
		On()
	})
	s.OnCancel(func(call *sip.Call) {
		log.Println("Cancel existing call")
		Off()
	})
	_ = <-forever
}

func GetOutboundIP(destination string) string {
	conn, err := net.Dial("udp", destination+":80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")

	return localAddr[0:idx]
}
