package main

import (
"log"
"net/http"
)

func main() {
// Create a new database
db, err := createDatabase()
if err != nil {
log.Fatal(err)
}

// Read the initial data from the database
initialClass, err := readData(db)
if err != nil {
log.Fatal(err)
}


// Write the updated data back to the database
err = db.Write("data", "class", initialClass)
if err != nil {
log.Fatal(err)
}

// Create a new router using Gorilla Mux
router := createRouter(initialClass)

// Start the server
log.Println("Server started on http://localhost:8080")
log.Fatal(http.ListenAndServe(":8080", router))
}
