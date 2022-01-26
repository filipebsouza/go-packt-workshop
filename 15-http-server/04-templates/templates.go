package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var tplStr = `
<html>
	<h1>Customer {{.ID}} 
	{{if .ID }} 
    <p>Details:</p> 
    <ul> 
      	{{if .Name}}<li>Name: {{.Name}}</li>{{end}} 
      	{{if .Surname}}<li>Surname: {{.Surname}}</li>{{end}} 
      	{{if .Age}}<li>Age: {{.Age}}</li>{{end}} 
   	</ul> 
   	{{else}} 
     	<p>Data not available</p> 
   	{{end}} 
</html>
`

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

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	vl := &urlValues{Values: r.URL.Query()}
	cust := Customer{}
	if vl.exist("id").Ok {
		cust.ID, _ = strconv.Atoi(vl.PropValue)
	}
	if vl.exist("name").Ok {
		cust.Name = vl.PropValue
	}
	if vl.exist("surname").Ok {
		cust.Surname = vl.PropValue
	}
	if vl.exist("age").Ok {
		cust.Age, _ = strconv.Atoi(vl.PropValue)
	}

	tmpl, _ := template.New("test").Parse(tplStr)
	tmpl.Execute(w, cust)
}

func main() {
	http.HandleFunc("/", HandleTemplate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
