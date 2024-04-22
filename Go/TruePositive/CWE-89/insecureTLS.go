package main

import (
    "crypto/tls"
    "fmt"
    "net/http"
)

func main() {
    // Vulnerable: Using insecure protocol and cipher suite
    config := &tls.Config{
        MinVersion: tls.VersionSSL30, // Insecure!
        CipherSuites: []uint16{
            tls.TLS_RSA_WITH_RC4_128_MD5, // Insecure!
        },
    }
    // ... (create server using insecure config)
}
