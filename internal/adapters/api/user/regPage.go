package user

import (
	"fmt"
	entities "forum/internal/model"
	"html/template"
	"log"
	"net/http"
)

// User Registration
func (h *Handler) RegPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RegPage")
	templ, err := template.ParseFiles("./view/templates/regPage.html")
	// if err != nil {
	// 	log.Println("HereHandlCreateUser", err)
	// 	w.WriteHeader(500)
	// 	// return
	// }

	switch r.Method {
	case "GET":
		input := &entities.User{}
		input.Err = false
		if r.URL.Path != "/regPage" {
			w.WriteHeader(404)
			return
		}
		if err = templ.Execute(w, input); err != nil {
			fmt.Println("here")
			w.WriteHeader(500)
			log.Println(err)
			return
		}
		return
	case "POST":
		fmt.Println(r.URL.Path, "#######")
		if r.URL.Path != "/regPage" {
			fmt.Println(r.URL.Path, "####################################")
			w.WriteHeader(404)
			return
		}
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(400)
			fmt.Println("1**************")
			return
		}
		fmt.Println("Len", len(r.Form), "**************")
		if len(r.Form) != 4 {
			log.Print("2************")
			w.WriteHeader(400)
			return
		}

		for v := range r.Form {
			if v != "uName" && v != "eMail" && v != "password1" && v != "password2" {
				log.Print("3************")
				w.WriteHeader(400)
				return
			}
		}

		user := entities.User{
			Name:      r.FormValue("uName"),
			Login:     r.FormValue("eMail"),
			Password1: r.FormValue("password1"),
			Password2: r.FormValue("password2"),
		}

		fmt.Println("CreateUser", user)
		if err := h.Service.CreateUser(user); err != nil {
			log.Print("4************")
			log.Println(err)
			w.WriteHeader(400)
			return
			user.Err = true
			user.MSG = err.Error()
			if err := templ.Execute(w, user); err != nil {
				w.WriteHeader(500)
				log.Println(err)
			}
			return
		}
	default:
		w.WriteHeader(405)
		return
	}
	fmt.Println("REDIRECT")
	// w.WriteHeader(200)
	// http.Redirect(w, r, "http://localhost:8080/logPage", 302)
}
