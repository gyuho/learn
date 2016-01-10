package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	privateKeyPath = "private.pem"
	certPath       = "cert.pem"
)

const port = ":8080"

func main() {
	certBytes, err := ioutil.ReadFile(certPath)
	if err != nil {
		panic(err)
	}

	clientCertPool := x509.NewCertPool()
	if ok := clientCertPool.AppendCertsFromPEM(certBytes); !ok {
		panic("can't add certificate to certificate pool!")
	}

	tlsConfigServer := &tls.Config{
		ClientAuth:               tls.RequireAndVerifyClientCert,
		ClientCAs:                clientCertPool,
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
	}
	tlsConfigServer.BuildNameToCertificate()

	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", handler)
	httpServer := &http.Server{
		Addr:      port,
		TLSConfig: tlsConfigServer,
		Handler:   mainRouter,
	}

	go func() {
		fmt.Println("Serving https://localhost" + port)

		// func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error
		if err := httpServer.ListenAndServeTLS(certPath, privateKeyPath); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	func() {
		fmt.Println("Sending client requests...")

		tlsCert, err := tls.LoadX509KeyPair(certPath, privateKeyPath)
		if err != nil {
			panic(err)
		}

		certBytes, err := ioutil.ReadFile(certPath)
		if err != nil {
			panic(err)
		}

		clientCertPool := x509.NewCertPool()
		if ok := clientCertPool.AppendCertsFromPEM(certBytes); !ok {
			panic("can't add certificate to certificate pool!")
		}

		tlsConfigClient := &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			RootCAs:      clientCertPool,
		}
		tlsConfigClient.BuildNameToCertificate()

		httpClient := http.DefaultClient
		httpClient.Transport = &http.Transport{TLSClientConfig: tlsConfigClient}
		resp, err := httpClient.Get("https://localhost" + port)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		rb, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("response:", string(rb))
	}()
}

func handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintln(w, "Hello World!")
		fmt.Fprintf(w, "req.TLS:        %+v\n", req.TLS)
		fmt.Fprintf(w, "DNSNames:       %#q\n", req.TLS.PeerCertificates[0].DNSNames)
		fmt.Fprintf(w, "EmailAddresses: %#q\n", req.TLS.PeerCertificates[0].EmailAddresses)
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

/*
Serving https://localhost:8080
Sending client requests...
response: Hello World!
req.TLS:        &{Version:771 HandshakeComplete:true DidResume:false CipherSuite:49199 NegotiatedProtocol: NegotiatedProtocolIsMutual:true ServerName:localhost PeerCertificates:[0xc82017e000] VerifiedChains:[[0xc82017e000 0xc8200a4000]] SignedCertificateTimestamps:[] OCSPResponse:[] TLSUnique:[156 55 205 75 79 27 21 192 84 244 36 226]}
DNSNames:       [`localhost`]
EmailAddresses: [`test@test.com`]
*/
