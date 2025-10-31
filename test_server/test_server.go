package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func testLinesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}
	fmt.Printf("[%s] Received %d bytes on /test-lines\n", time.Now().Format("15:04:05"), len(body))
	fmt.Printf("Content-Type: %s\n", r.Header.Get("Content-Type"))
	fmt.Printf("Headers: %+v\n", r.Header)
	lines := string(body)
	if len(lines) > 500 {
		fmt.Printf("First 500 chars: %s...\n", lines[:500])
	} else {
		fmt.Printf("Content: %s\n", lines)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/test-lines", testLinesHandler)
	fmt.Println("Test server running on :8080")
	fmt.Println("Endpoint: http://localhost:8080/test-lines")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
