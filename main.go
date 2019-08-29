package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
)

func GetIP() net.IP {
	conn, _ := net.Dial("udp", "1.2.3.4:80")
	conn.Close()
	ip := conn.LocalAddr().(*net.UDPAddr)
	return ip.IP
}

func main() {
	// Collect memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// Collect hostname
	h, _ := os.Hostname()

	// Handler for "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Print collected info with newline
		fmt.Fprintf(w, "Hostname: %s\n", h)
		fmt.Fprintf(w, "IP: %s\n", GetIP())
		fmt.Fprintf(w, "Platform: %s\n", runtime.GOOS)
		fmt.Fprintf(w, "CPU: %d\n", runtime.NumCPU())
		fmt.Fprintf(w, "RAM: %d", m.Sys)
	})

	// Start the web server
	log.Fatal(http.ListenAndServe(":8082", nil))
}
