package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/levigross/grequests"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert.pem", "private.pem")
	if err != nil {
		log.Fatalln("Unable to load cert", err)
	}

	clientCACert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal("Unable to open cert", err)
	}

	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      clientCertPool,
	}

	tlsConfig.BuildNameToCertificate()
	ro := &grequests.RequestOptions{
		HTTPClient: &http.Client{
			Transport: &http.Transport{TLSClientConfig: tlsConfig},
		},
	}
	resp, err := grequests.Get("https://localhost:8080", ro)
	if err != nil {
		log.Println("Unable to speak to our server", err)
	}

	log.Println(resp.String())
}
