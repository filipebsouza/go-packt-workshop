package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Customer struct {
	ID, Age       int
	Name, Surname string
}

type urlValues struct {
	url.Values
	Ok        bool
	PropValue string
}

func (u *urlValues) exist(key string) *urlValues {
	vl, ok := u.Values[key]
	if ok {
		u.PropValue = strings.Join(vl, ",")
	}
	u.Ok = ok

	return u
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	vl := &urlValues{Values: r.URL.Query()}
	cust := Customer{}
	if vl.exist("name").Ok {
		cust.Name = vl.PropValue
	}

	tmpl, _ := template.ParseFiles("./pages/home.html")
	tmpl.Execute(w, cust)
}

func main() {
	http.HandleFunc("/", HandleHome)
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "./pages/form.html")
			return
		}

		if r.Method == http.MethodPost {
			cust := Customer{}
			err := r.ParseForm()
			if err != nil {
				w.WriteHeader(400)
				return
			}

			cust.Name = r.Form.Get("name")

			cust.Surname = r.Form.Get("surname")
			cust.Age, _ = strconv.Atoi(r.Form.Get("age"))

			tmpl, _ := template.ParseFiles("./pages/form.html")
			tmpl.Execute(w, cust)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
