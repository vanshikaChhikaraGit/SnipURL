package main

import (
	"URL_SHORTNER/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)
var ctx = context.Background()
func main(){
	err:=godotenv.Load()
	if err!=nil{
		fmt.Print("error loading env variables")
		return;
	}
	BASE_URL:= os.Getenv("BASE_URL")

	rdb:=utils.RedisClient()
	if rdb==nil{
		fmt.Println("Failed to connect to Redis")
        return
	}
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request)  {
		tmp:=template.Must(template.ParseFiles("tempalate/index.html"))
		tmp.Execute(w,nil)
	})

	http.HandleFunc("/shortenURL",func(w http.ResponseWriter, r *http.Request) {
        url:= r.FormValue("url")
		fmt.Println("input url=",url)
		shortURL:= utils.ShortenURL(url)
		fmt.Println("short code for the given url is ",shortURL)
		fullShortURL:= fmt.Sprintf("%s/r/%s",BASE_URL,shortURL)

		fmt.Println("shortened url is",fullShortURL)

		utils.SetKey(&ctx,rdb,shortURL,url)
		
		fmt.Fprintf(w,`<p class="mt-4 text-green-600">Shortened URL: <a 
href="/r/%s" class="underline">%s</a></p>`,shortURL,fullShortURL)
	})

	http.HandleFunc("/r/",func(w http.ResponseWriter, r *http.Request) {
		pathURL:= r.URL.Path
		shortCode:= pathURL[len("/r/"):]
		fmt.Println(shortCode)
		if shortCode==""{
			http.Error(w, "Invalid URL", http.StatusBadRequest) 
			return 
		}
		longURL,err:= utils.GetLongUrl(&ctx,rdb,shortCode)
		if err!=nil{
			http.Error(w, "URL not found", http.StatusNotFound) 
			return 
		}

		http.Redirect(w,r,longURL,http.StatusPermanentRedirect)

	})
	// Get the PORT from environment variables (Render sets it automatically)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if no PORT environment variable is set (useful for local development)
	}

	// Start the server on the correct port
	fmt.Printf("Server is running on port %s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}