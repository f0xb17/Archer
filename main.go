package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// The URL to the Arch Linux packages search API
const repoURL = "https://archlinux.org/packages/search/json/?q="

// Package represents a single package in the search results
type Package struct {
	Name        string `json:"pkgname"` // The name of the package
	Description string `json:"pkgdesc"` // The description of the package
}

// SearchResult represents the search results from the Arch Linux packages API
type SearchResult struct {
	Packages []Package `json:"results"`
}

// searchRepo queries the Arch Linux packages API with the given search term
// and returns a slice of packages that match the search term
func searchRepo(query string) ([]Package, error) {
	response, err := http.Get(repoURL + query)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", response.StatusCode)
	}

	var result SearchResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Packages, nil
}

// The main function is the entry point of the program
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: archer <query>")
		os.Exit(1)
	}

	query := os.Args[1]
	packages, err := searchRepo(query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	for _, pkg := range packages {
		fmt.Printf("%s - %s\n", pkg.Name, pkg.Description)
		fmt.Println()
	}
}
