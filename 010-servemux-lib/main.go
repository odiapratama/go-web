package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	router := httprouter.New()

	router.GET("/", index)
	router.GET("/user/:name", user)
	router.GET("/blog/:category/:article", blogRead)
	router.POST("/blog/:category/:article", blogWrite)
	router.GET("/about", about)
	router.GET("/contact", contact)
	router.GET("/apply", apply)
	router.POST("/apply", applyProcess)
	router.GET("/redirect", redirect)
	router.GET("/redirected", redirected)
	router.GET("/panic", panicPage)
	router.GET("/recover", recoverPage)
	router.GET("/notfound", customNotFound)
	router.NotFound = http.HandlerFunc(customNotFoundHandler)

	http.Handle("/", router)

	http.HandleFunc("/500", internalServerError)
	http.HandleFunc("/405", methodNotAllowed)
	http.HandleFunc("/400", badRequest)
	http.HandleFunc("/401", unauthorized)
	http.HandleFunc("/403", forbidden)

	http.ListenAndServe(":8080", nil)
}

func user(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "USER, %s!\n", p.ByName("name"))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find the page you were looking for :(</h1>")
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "<h1>Internal Server Error</h1>")
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprint(w, "<h1>Method Not Allowed</h1>")
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "<h1>Bad Request</h1>")
}

func unauthorized(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, "<h1>Unauthorized</h1>")
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprint(w, "<h1>Forbidden</h1>")
}

func redirect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/redirected", http.StatusSeeOther)
}

func redirected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "You have been redirected")
}

func panicPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("PANIC")
}

func recoverPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	}()
	panic("PANIC")
}

func customNotFound(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.NotFound(w, r)
}

func customNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Custom Not Found</h1>")
}

func blogRead(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "READ CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(w, "READ ARTICLE, %s!\n", ps.ByName("article"))
}

func blogWrite(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "WRITE CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(w, "WRITE ARTICLE, %s!\n", ps.ByName("article"))
}

func about(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
