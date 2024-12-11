package main

import (
	"URL_SHORTNER/utils"
	"context"
	"fmt"
	"net/http"
	"os"
)

// Global port variable
var port string

var ctx = context.Background()

func corsMiddleware(next http.Handler)http.Handler{
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
w.Header().Set("Access-Control-Allow-Origin", "https://your-frontend.vercel.app")
w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// Preflight request
if r.Method == http.MethodOptions {
	w.WriteHeader(http.StatusNoContent)
	return
}

next.ServeHTTP(w, r)
})
}

// Global function to initialize the application settings
func init() {
	
// Get the BASE_URL from the environment
	BASE_URL := os.Getenv("BASE_URL")

	// Connect to Redis (ensure that Redis is properly configured in the environment variables)
	rdb := utils.RedisClient()
	if rdb == nil {
		fmt.Println("Failed to connect to Redis")
		return
	}

	// Setup routes
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	tmp := template.Must(template.ParseFiles("client/index.html"))
	// 	tmp.Execute(w, nil)
	// })

	http.HandleFunc("/shortenURL", func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("url")
		fmt.Println("Input URL:", url)
		shortURL := utils.ShortenURL(url)
		fmt.Println("Short code for the given URL is:", shortURL)
		fullShortURL := fmt.Sprintf("%s/r/%s", BASE_URL, shortURL)

		fmt.Println("Shortened URL is:", fullShortURL)

		utils.SetKey(&ctx, rdb, shortURL, url)

		// Set content type to JSON and return the response
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"shortUrl": "%s"}`, fullShortURL)
	})

	http.HandleFunc("/r/", func(w http.ResponseWriter, r *http.Request) {
		pathURL := r.URL.Path
		shortCode := pathURL[len("/r/"):]
		fmt.Println(shortCode)
		if shortCode == "" {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
		longURL, err := utils.GetLongUrl(&ctx, rdb, shortCode)
		if err != nil {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, longURL, http.StatusPermanentRedirect)
	})

	// Get the PORT from environment variables (Render sets it automatically)
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if no PORT environment variable is set (useful for local development)
	}
}

// Main function to start the server
func main() {
	// Start the server on the correct port
	fmt.Printf("Server is running on port %s\n", port)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
