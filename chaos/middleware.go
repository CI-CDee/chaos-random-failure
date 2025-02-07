package chaos

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Middleware introduces chaos (latency, errors, crashes) into the system
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Apply chaos based on probabilities
		applyLatency()
		applyError(w)
		applyCrash()

		next.ServeHTTP(w, r)
	})
}

func applyLatency() {
	latencyStr := os.Getenv("CHAOS_LATENCY_MS")
	if latencyStr == "" {
		return
	}
	latency, err := strconv.Atoi(latencyStr)
	if err != nil {
		return
	}
	time.Sleep(time.Duration(latency) * time.Millisecond)
}

func applyError(w http.ResponseWriter) {
	errorRateStr := os.Getenv("CHAOS_ERROR_RATE")
	if errorRateStr == "" {
		return
	}
	errorRate, err := strconv.ParseFloat(errorRateStr, 64)
	if err != nil {
		return
	}
	if rand.Float64() < errorRate {
		log.Println("Chaos: Injecting error")
		http.Error(w, "Internal Server Error (Chaos)", http.StatusInternalServerError)
	}
}

func applyCrash() {
	crashRateStr := os.Getenv("CHAOS_CRASH_RATE")
	if crashRateStr == "" {
		return
	}
	crashRate, err := strconv.ParseFloat(crashRateStr, 64)
	if err != nil {
		return
	}
	if rand.Float64() < crashRate {
		log.Println("Chaos: Simulating server crash")
		os.Exit(1) // Crash the server
	}
}
