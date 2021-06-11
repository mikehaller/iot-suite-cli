package iotsuite

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/fatih/color"
	"net/http/httptrace"
	"crypto/tls"
)

func Get(httpClient *http.Client, url string) {
	color.Cyan(url)
	req,err := http.NewRequest(http.MethodGet, url, nil)
	if (err != nil) {
			log.Fatal("Unable to create HTTP NewRequest",err)
	}
	resp,err := httpClient.Do(req)
	if (err != nil) {
			log.Fatal("General HTTP Client:",err)
	}
	defer resp.Body.Close()
	DumpJsonResponse(resp)
}

func newHttpTrace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			log.WithFields(log.Fields{"hostPort": hostPort}).Trace("GetConn")
		},
        GotConn: func(connInfo httptrace.GotConnInfo) {
			log.WithFields(log.Fields{
					"LocalAddr": connInfo.Conn.LocalAddr(),
					"RemoteAddr": connInfo.Conn.RemoteAddr(),
					"Reused": connInfo.Reused,
					"WasIdle": connInfo.WasIdle,
					"IdleTime": connInfo.IdleTime}).Trace("GotConn")
        },
        DNSStart: func(dnsInfo httptrace.DNSStartInfo) {
        	log.WithFields(log.Fields{"dnsInfo": dnsInfo}).Trace("DNSStart")
        },
        DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
        	log.WithFields(log.Fields{"dnsInfo": dnsInfo}).Trace("DNSDone")
        },
        PutIdleConn: func(err error) {
            log.WithFields(log.Fields{"err": err}).Trace("PutIdleConn")
        },
        GotFirstResponseByte : func() {
        	log.Trace("GotFirstResponseByte")
        },
        Got100Continue: func() {
        	log.Trace("HTTP Got 100 Continue'\n")
        },
        ConnectStart: func(network, addr string) {
        	log.WithFields(log.Fields{"network": network, "addr":addr}).Trace("ConnectStart")
        },
        ConnectDone: func(network, addr string, err error) {
        	log.WithFields(log.Fields{"network": network, "addr":addr,"err":err}).Trace("ConnectDone")
        },
        TLSHandshakeStart: func() {
        	log.Trace("TLSHandshakeStart")
        },
        TLSHandshakeDone: func(tlsState tls.ConnectionState, err error) {
        	log.WithFields(log.Fields{"tlsState": tlsState, "err":err}).Trace("TLSHandshakeDone")
        },
		WroteHeaderField: func(key string, value []string) {
        	log.WithFields(log.Fields{"key": key, "value":value}).Trace("WroteHeaderField")
		},
		WroteHeaders: func() {
        	log.Trace("WroteHeaders")
		},
		Wait100Continue: func() {
        	log.Trace("Wait100Continue")
		},
		WroteRequest: func(requestInfo httptrace.WroteRequestInfo) {
			log.WithFields(log.Fields{"requestInfo": requestInfo}).Trace("WroteRequest")
		},
    }
}