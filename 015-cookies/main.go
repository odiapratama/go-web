package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/get", getCookie)
	http.HandleFunc("/delete", deleteCookie)
	http.HandleFunc("/update", updateCookie)
	http.HandleFunc("/expiration", setCookieWithExpiration)
	http.HandleFunc("/secure", setCookieWithSecure)
	http.HandleFunc("/http-only", setCookieWithHttpOnly)
	http.HandleFunc("/same-site", setCookieWithSameSite)
	http.HandleFunc("/with-domain", setCookieWithDomain)

	http.ListenAndServe(":8080", nil)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "my-value",
	})

	fmt.Fprintln(w, "Cookie created")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, "Cookie not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Cookie found: %s\n", cookie.Value)
}

func deleteCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, "Cookie not found", http.StatusNotFound)
		return
	}

	cookie.MaxAge = -1
	http.SetCookie(w, cookie)

	fmt.Fprintln(w, "Cookie deleted")
}

func updateCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, "Cookie not found", http.StatusNotFound)
		return
	}

	cookie.Value = "new-value"
	http.SetCookie(w, cookie)

	fmt.Fprintln(w, "Cookie updated")
}

func setCookieWithExpiration(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "my-cookie",
		Value:   "my-value",
		Expires: time.Now().Add(24 * time.Hour),
	})

	fmt.Fprintln(w, "Cookie created")
}

func setCookieWithSecure(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "my-cookie",
		Value:  "my-value",
		Secure: true,
	})

	fmt.Fprintln(w, "Cookie created")
}

func setCookieWithHttpOnly(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "my-cookie",
		Value:    "my-value",
		HttpOnly: true,
	})

	fmt.Fprintln(w, "Cookie created")
}

func setCookieWithSameSite(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "my-cookie",
		Value:    "my-value",
		SameSite: http.SameSiteStrictMode,
	})

	fmt.Fprintln(w, "Cookie created")
}

func setCookieWithDomain(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "my-cookie",
		Value:  "my-value",
		Domain: "example.com",
	})

	fmt.Fprintln(w, "Cookie created")
}
