package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

const formHTML = `<html>
<head>
<title></title>
</head>
<body>
<form action="/senddatahere" method="post">
    Username:<input type="text" name="username">
    Password:<input type="password" name="password">
    <input type="submit" value="Login">
</form>
</body>
</html>
`

var formTemplate = template.Must(template.New("form").Parse(formHTML))

var names = []string{
	"George",
	"Suneil",
}

func handler(w http.ResponseWriter, r *http.Request) {
	for i, name := range names {
		fmt.Fprintf(w, "%d: %s\n", i, name)
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	urlBits := strings.Split(r.URL.Path, "/")
	name := urlBits[len(urlBits)-1]

	names = append(names, name)
	fmt.Fprintf(w, "%d", len(names))
}

func loginForm(w http.ResponseWriter, r *http.Request) {
	formTemplate.Execute(w, nil)
}

func sendDataHere(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "username: %s", r.Form["username"])
	fmt.Fprintf(w, "password: %s", r.Form["password"])
}

func main() {
	fmt.Println("Trying to run")
	http.HandleFunc("/senddatahere/", sendDataHere)
	http.HandleFunc("/login/", loginForm)
	http.HandleFunc("/add/", addHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
