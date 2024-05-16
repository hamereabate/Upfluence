package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func AnalysisHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	duration := r.URL.Query().Get("duration")
	dimension := r.URL.Query().Get("dimension")

	d, err := time.ParseDuration(duration)
	if err != nil {
		http.Error(w, "Invalid duration", http.StatusBadRequest)
		return
	}

	endTime := time.Now().Add(d)

	var totalPosts int
	var minTimestamp, maxTimestamp int64
	var totalLikes int

	resp, err := http.Get("https://stream.upfluence.co/stream")
	if err != nil {
		http.Error(w, "Error connecting to Upfluence API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	dimensionRegex := getDimensionRegex(dimension)
	timStampRegex := getTimestampRegex()

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data:") {
			continue
		}

		data := strings.TrimPrefix(line, "data:")
		if !strings.Contains(data, dimension) {
			continue
		}

		dimensionMatches := dimensionRegex.FindAllStringSubmatch(data, -1)
		timeStampMatches := timStampRegex.FindAllStringSubmatch(data, -1)

		if len(dimensionMatches) == 0 {
			continue
		}

		if len(dimensionMatches[0]) < 2 {
			fmt.Println("Uxpected dimension matches:", len(dimensionMatches[0]))
			continue
		}
		likes, err := strconv.Atoi(dimensionMatches[0][1])
		if err != nil {
			fmt.Println("Error converting likes value:", err)
			continue
		}
		timestamp, err := strconv.ParseInt(timeStampMatches[0][1], 10, 64)
		if err != nil {
			fmt.Println("Error converting timestamp:", err)
			continue
		}

		totalPosts++
		totalLikes += likes

		if minTimestamp == 0 || timestamp < minTimestamp {
			minTimestamp = timestamp
		}
		if maxTimestamp == 0 || timestamp > maxTimestamp {
			maxTimestamp = timestamp
		}

		if time.Now().After(endTime) {
			break
		}
	}

	avgLikes := 0
	if totalPosts > 0 {
		avgLikes = totalLikes / totalPosts
	}

	response := Response{
		TotalPosts:       totalPosts,
		MinimumTimestamp: minTimestamp,
		MaximumTimestamp: maxTimestamp,
		AvgLikes:         avgLikes,
	}

	// Set content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode response as JSON and send to client
	json.NewEncoder(w).Encode(response)

}

func getDimensionRegex(dimension string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`"%s":(\d+)`, dimension))
}

func getTimestampRegex() *regexp.Regexp {
	return regexp.MustCompile(`"timestamp":(\d+)`)
}
