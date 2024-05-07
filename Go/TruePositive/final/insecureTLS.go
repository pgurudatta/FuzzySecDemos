package main

import (
    "crypto/tls"
    "fmt"
    "net/http"
)

func main() {
    // Create a secure TLS configuration
    config := &tls.Config{
        MinVersion: tls.VersionTLS12, // Set the minimum TLS version to TLS 1.2
        // Use recommended secure cipher suites
        CipherSuites: []uint16{
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
            tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
            tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
        },
    }

    // Create an HTTP server with the secure TLS configuration
    server := &http.Server{
        Addr:      ":443",
        Handler:   nil, // Use default handler
        TLSConfig: config,
    }

    // Start the server
    err := server.ListenAndServeTLS("cert.pem", "key.pem")
    if err != nil {
        fmt.Printf("Failed to start server: %s\n", err)
        return
    }
}
