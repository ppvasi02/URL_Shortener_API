package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	/*"database/sql"
	"encoding/json"
	"fmt"
	"io"


	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"*/)

type URL struct {
	LongURL   string `json:"long_url"`
	ShortCode string `json:"short_code"`
}

var links = []URL{
	{LongURL: "https://www.google.com/", ShortCode: "000001"},
	{LongURL: "https://www.youtube.com/", ShortCode: "000002"},
}

/*func generateShortCode() (string, error) {
	// uuid (fixed length - can be modified)
	shortCode := uuid.New().String()[:6] // Get first 6 characters from UUID
	return shortCode, nil
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var urlData URL
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&urlData)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Generate short code
	shortCode, err := generateShortCode()
	if err != nil {
		fmt.Println("Error generating short code:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ASSUMES DATABASE CONNECTION ALREADY PRESENT
	// Database interaction (replace with your specific database driver)
	db, err := connectToDatabase() // Connect to the database
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close() // Close the database connection after use

	// Save the mapping (replace with actual insert statement)
	stmt, err := db.Prepare("INSERT INTO urls (long_url, short_code) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Error preparing insert statement:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(urlData.LongURL, shortCode)
	if err != nil {
		fmt.Println("Error saving URL mapping:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	stmt.Close() // Close the prepared statement

	// Respond with JSON containing the shortened URL
	response := URL{LongURL: urlData.LongURL, ShortCode: shortCode}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		fmt.Println("Error encoding response:", err)
		return
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// Extract short code from the request path
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// Database interaction (replace with your specific database driver)
	db, err := connectToDatabase() // Connect to the database
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close() // Close the database connection after use

	// Retrieve the long URL (replace with actual query statement)
	var longURL string
	stmt, err := db.Prepare("SELECT long_url FROM urls WHERE short_code = ?")
	if err != nil {
		fmt.Println("Error preparing select statement:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = stmt.QueryRow(shortCode).Scan(&longURL)
	if err != nil {
		if err == sql.ErrNoRows { // Handle case where short code is not found
			fmt.Println("Short code not found:", shortCode)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Println("Error retrieving long URL:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	stmt.Close() // Close the prepared statement

	// Redirect to the long URL
	http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
}

/* func (client *mongo.Client) Close(ctx context.Context) error {
    // This is the actual method for closing a MongoDB connection using the official Go driver
    return client.Disconnect(ctx)
}
*/

/*func (db *DB) Prepare(query string) (*Stmt, error) {
    // This is the `Prepare` function from the standard library `database/sql` package
    return db.conn.PrepareContext(ctx, query)
}
*/

/*// Function to connect to the MongoDB database
func connectToDatabase() (*mongo.Client, error) {
    // Replace with your actual connection details
    uri := "mongodb://localhost:27017"
    dbName := "url_shortener"

    ctx := context.Background()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, fmt.Errorf("error pinging MongoDB: %w", err)
    }

    fmt.Println("Connected to MongoDB")
    return client, nil
} */
/*
func connectToDatabase() (*http.Client, error) {
	// Use HTTP client to call the Node.js script for database connection
	url := "http://localhost:3000/connect" // /Users/ppvas/Desktop/API/database.js  http://localhost:PORT/connect
	client := &http.Client{}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request to database.js: %w", err)
	}

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to database.js: %w", err)
	}
	defer resp.Body.Close()

	// Check for successful response status code
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body from database.js: %w", err)
		}
		return nil, fmt.Errorf("error connecting to database: %s", string(body))
	}

	// Decode the response body (assuming JSON format)
	var connectionInfo struct {
		MongoClient string `json:"mongoClient"` // Replace with actual field names if different
	}
	err = json.NewDecoder(resp.Body).Decode(&connectionInfo)
	if err != nil {
		return nil, fmt.Errorf("error decoding response from database.js: %w", err)
	}

	// Handle the returned connection information (replace with your actual logic)
	fmt.Println("Connected to database using information from database.js")
	// You might need to parse the connectionInfo and use it to establish a Go driver connection

	return client, nil
}*/

func getLinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, links)
}

func createLink(c *gin.Context) {
	var newLink URL

	if err := c.BindJSON(&newLink); err != nil {
		return
	}

	links = append(links, newLink)
	c.IndentedJSON(http.StatusCreated, newLink)
}

func requestLink(c *gin.Context) {
	code := c.Param("short_code")
	URL, err := getLink(code)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Link not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, URL)
}

func getLink(code string) (*URL, error) {
	for i, l := range links {
		if l.ShortCode == code {
			return &links[i], nil
		}
	}
	return nil, errors.New("code not found")
}

/*func main() {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", shortenURLHandler).Methods("POST")
	router.HandleFunc("/{shortCode}", redirectHandler).Methods("GET")
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", router)
}*/

func main() {
	router := gin.Default()
	router.GET("/links", getLinks)
	router.GET("/links/:short_code", requestLink)
	router.POST("/links", createLink)
	router.Run("localhost:8080")
}
