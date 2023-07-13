package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
)

// Define a struct to represent the user data
type User struct {
	Login       string `json:"login"`
	Repositories int    `json:"repositoryCount"`
}
// UserService
func main() {
	// ...

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
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

