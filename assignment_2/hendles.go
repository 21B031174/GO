package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("templates/register.html")
	tmpl.Execute(w, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("templates/home.html")
	tmpl.Execute(w, nil)
}

func checking(w http.ResponseWriter, r *http.Request) {
	settings := ReadJStr()

	if r.FormValue("type") == "Login" {
		for _, user := range settings.Users {
			if user.Username == r.FormValue("login") && ToHash(r.FormValue("password")) == user.Password {
				http.Redirect(w, r, "/home/", http.StatusSeeOther)
			}
		}

	} else {
		tmpl := template.New("data")
		settings := ReadJStr()
		for _, user := range settings.Users {
			if user.Username == r.FormValue("login") {
				message := "Successfully registered."
				tmpl.Parse(`<h1>{{.}}</h1>`)
				tmpl.Execute(w, message)
				return
			}
		}
		newUser := User{r.FormValue("login"), ToHash(r.FormValue("password"))}
		settings.Users = append(settings.Users, newUser)
		settings.WriteJStr()
		message := "Successfully registered."
		tmpl.Parse(`<h1>{{.}}</h1>`)
		tmpl.Execute(w, message)
	}
}
