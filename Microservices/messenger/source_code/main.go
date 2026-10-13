package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

var startTime = time.Now()

// links holds optional public URLs of the other services for the home page menu,
// e.g. AUTH_URL=https://auth.example.com
func links() map[string]string {
	out := map[string]string{}
	for _, s := range []string{"auth", "profile", "messenger", "finance", "share"} {
		if v := os.Getenv(strings.ToUpper(s) + "_URL"); v != "" {
			out[s] = v
		}
	}
	return out
}

func homeHandler(page []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(page)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"service":       "messenger",
		"status":        "UP",
		"runtime":       "Go " + runtime.Version()[2:],
		"hostname":      hostname,
		"uptimeSeconds": int(time.Since(startTime).Seconds()),
		"timestamp":     time.Now().UTC().Format(time.RFC3339),
		"links":         links(),
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Hello from Go Messenger Service")
}

func main() {
	page, err := os.ReadFile("static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", homeHandler(page))
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":3002", nil))
}
