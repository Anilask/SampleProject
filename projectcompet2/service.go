package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
)

// Define a struct to represent the repository details
type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"html_url"`
}

// RepositoryService
func main() {
	// ...

	http.HandleFunc("/repository", func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.github.com/repos/USERNAME/REPO", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Authorization header with the personal access token
		req.Header.Set("Authorization", "Bearer YOUR_PERSONAL_ACCESS_TOKEN")

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// ...
	})

	// ...
}