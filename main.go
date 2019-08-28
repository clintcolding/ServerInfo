package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

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
		fmt.Fprintf(w, "Platform: %s\n", runtime.GOOS)
		fmt.Fprintf(w, "CPU: %d\n", runtime.NumCPU())
		fmt.Fprintf(w, "RAM: %d", m.Sys)
	})

	// Start the web server
	log.Fatal(http.ListenAndServe(":8082", nil))
}
