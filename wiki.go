package main

import (
	"fmt"
	"io/ioutil"
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

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
