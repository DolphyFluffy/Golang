package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Describes how page data will be stored in memory
type Page struct {
	Title string
	Body  []byte // This type means a byte slice - this si the type expected by the io libraries
}

// save() is a method w/o params that takes a reciever p (pointer to a page) and returns a value of type error
// Saves the page's body to a text file
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) // WriteFile returns an error or nil (if no error)
} // 0600 means the file has read/write permissions for currnet user only

// loadPage cnstructs the file name from the title, reads file contents into the body and returns a pointer
// to a page literal with the title and body
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// viewHandler allows ysers to view a wiki page
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):] // Extracts the URL after the static '/view/' path
	p, _ := loadPage(title)             // Loads page with the title in the url and writes to w http.ResponseWriter
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)       // Tells http pacakge to handle reqquest with '/view/' root
	log.Fatal(http.ListenAndServe(":8080", nil)) // Listens on port 8080 until an error occurs or termination
}
