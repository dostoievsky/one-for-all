package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"strconv"
	"strings"
	)
	
	func createRouter(initialClass []Class) *mux.Router {
	router := mux.NewRouter()
	
	router.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body
		var requestData struct {
		Email string `json:"email"`
		Noun  string `json:"noun"`
		}
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
		}
		
		// Check if the email and noun exist in the database
		exists := false
		for _, item := range initialClass {
		if item.Email == requestData.Email && item.Noun == requestData.Noun {
		exists = true
		break
		}
		}
		
		// Prepare the response
		responseData := struct {
		Exists bool `json:"exists"`
		}{
		Exists: exists,
		}
		
		// Send the response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(responseData)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		
		// Log the request and response
		log.Printf("Request: %+v\n", requestData)
		log.Printf("Response: %+v\n", responseData)
		}).Methods("POST")

		router.HandleFunc("/votes", func(w http.ResponseWriter, r *http.Request) {
			// Count the number of true votes
			votedCount := 0
			for _, item := range initialClass {
			if item.Voted {
			votedCount++
			}
			}
			
			// Prepare the response
			responseData := struct {
			Result string `json:"result"`
			}{
			Result: fmt.Sprintf("%d/%d", votedCount, len(initialClass)),
			}
			
			// Send the response
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(responseData)
			if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
			}
		}).Methods("GET")		


		router.HandleFunc("/vote", func(w http.ResponseWriter, r *http.Request) {
			// Parse the request body
			var requestData struct {
			Email string   `json:"email"`
			Noun  string   `json:"noun"`
			Votes []string `json:"votes"`
			Time  string   `json:"time"`
			}
			err := json.NewDecoder(r.Body).Decode(&requestData)
			if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
			}
			
			// Convert votes to integers
			intVotes := make([]int, len(requestData.Votes))
			for i, vote := range requestData.Votes {
			intVote, err := strconv.Atoi(vote)

			if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
			}
			intVotes[i] = intVote
			}
			timeVoted := strings.Split(requestData.Time, ",")

			// Find the corresponding item in initialClass
			for i, item := range initialClass {
			if item.Email == requestData.Email && item.Noun == requestData.Noun {
			// Update the item in initialClass
			initialClass[i].Vote = intVotes
			initialClass[i].TimeVoted = timeVoted
			initialClass[i].Voted = true
			break
			}
			}
			
			// Prepare the response
			responseData := struct {
			Success bool `json:"success"`
			}{
			Success: true,
			}
			
			// Send the response
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(responseData)
			if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
			}
			
			// Log the request and response
			log.Printf("Request: %+v\n", requestData)
			log.Printf("Response: %+v\n", responseData)
		}).Methods("POST")
				
		router.HandleFunc("/eachVoted", func(w http.ResponseWriter, r *http.Request) {
			// Create an array to store the vote counts
			voteCounts := make([]int, 6) // Assuming the vote array can have values from 1 to 6
			
			// Count the occurrences of each number in the Vote array
			for _, item := range initialClass {
			if item.Voted {
			for _, vote := range item.Vote {
			if vote >= 1 && vote <= 6 {
			voteCounts[vote-1]++
			}
			}
			}
			}
			
			// Prepare the response
			responseData := struct {
			VoteCounts []int `json:"voteCounts"`
			}{
			VoteCounts: voteCounts,
			}
			
			// Send the response
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(responseData)
			if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
			}
			}).Methods("GET")
			
	
// Route for the root URL ("/")
router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
http.ServeFile(w, r, "index.html")
})

return router
}
