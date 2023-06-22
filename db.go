package main

import (
"log"

"github.com/nanobox-io/golang-scribble"
)

type Class struct {
	Email     string   `json:"email"`
	Noun      string   `json:"noun"`
	Vote      []int    `json:"vote"`
	Voted 	  bool     `json:"voted"`
	TimeVoted []string `json:"timeVoted"`
}

func createDatabase() (*scribble.Driver, error) {
	// Create a new Scribble database
	db, err := scribble.New(".", nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func readData(db *scribble.Driver) ([]Class, error) {
	// Read the initial class.json file
	var initialClass []Class
	err := db.Read("data", "class", &initialClass)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	
	// Check if the fields are missing and initialize them
	for i := range initialClass {
		if initialClass[i].Vote == nil {
			initialClass[i].Vote = []int{}
		}
		if initialClass[i].TimeVoted == nil {
			initialClass[i].TimeVoted = []string{}
		}
		if !initialClass[i].Voted {
			initialClass[i].Voted = false
		}
	}
	
	return initialClass, nil
}
	
	
func writeData(db *scribble.Driver, data []Class) error {
	// Write the updated data back to the database
	err := db.Write("data", "class", data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
	