package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"net/url"
)

const MOVIEURL = "https://omdbapi.com/"

type MOVIE struct{
	Title string
	Year string
	Rated string
}

func main(){

	searchString := url.QueryEscape("star wars")
	resp, err := http.Get(MOVIEURL + "?t=" + searchString)
	if err != nil {
		fmt.Errorf("something failed: %s", err)
	}

	//We must close resp.Body on all execution paths
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		fmt.Errorf("search query failed: %s", resp.Status)
	}

	var movie MOVIE
	error := json.NewDecoder(resp.Body).Decode(&movie)
	if error != nil {
		resp.Body.Close()
		fmt.Errorf("error on decode: %s", error)
	}

	resp.Body.Close()
	fmt.Println(movie)
	fmt.Printf("Movie Title: %s, Year Published: %s, Rating: %s", movie.Title, movie.Year, movie.Rated)

}