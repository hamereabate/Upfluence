package main

import (
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	TotalPosts       int   `json:"total_posts"`
	MinimumTimestamp int64 `json:"minimum_timestamp"`
	MaximumTimestamp int64 `json:"maximum_timestamp"`
	AvgLikes         int   `json:"avg_likes"`
}

func main() {
	//analysis := AnalysisHandler()
	http.HandleFunc("/analysis", AnalysisHandler)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
