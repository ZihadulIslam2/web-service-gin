package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))

    p3 := &Page{Title: "test", Body: []byte("This is a test Page.")}
    p3.save()

      http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}


type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return  os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/") :]
    p, err := loadPage(title)
    if err != nil {
        http.Error(w, "Page not found", http.StatusNotFound)
        return
    }
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}