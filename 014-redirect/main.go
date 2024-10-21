package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/see", seeOther)
	http.HandleFunc("/temporary", temporaryDirect)
	http.HandleFunc("/moved", movedPermanently)
	http.HandleFunc("/barred", barred)
	http.HandleFunc("/write", writeHeader)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}

	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func seeOther(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at seeOther:", req.Method)
	// process form submission here
	http.Redirect(w, req, "/barred", http.StatusSeeOther)
}

func temporaryDirect(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at temporaryDirect:", req.Method)
	// process form submission here
	http.Redirect(w, req, "/barred", http.StatusTemporaryRedirect)
}

func movedPermanently(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at movedPermanently:", req.Method)
	http.Redirect(w, req, "/barred", http.StatusMovedPermanently)
}

func writeHeader(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at writeHeader:", req.Method)
	// process form data
	w.Header().Set("Location", "/barred")
	w.WriteHeader(http.StatusSeeOther)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
