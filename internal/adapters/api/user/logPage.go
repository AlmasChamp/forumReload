package user

import (
	"fmt"
	entities "forum/internal/model"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func (h *Handler) logPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logPage", r.Method)
	templ, err := template.ParseFiles("./view/templates/signIn.html")
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
	user := entities.User{}

	switch r.Method {
	case "GET":
		if r.URL.Path != "/logPage" {
			w.WriteHeader(404)
			return
		}
		// er := entities.Error{}
		if err = templ.Execute(w, user); err != nil {
			user.Err = true
			user.MSG = err.Error()
			fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^")
			w.WriteHeader(500)
			log.Println(err)
			return
		}
	case "POST":
		if r.URL.Path != "/logPage" {
			w.WriteHeader(404)
			return
		}
		cookie := &http.Cookie{
			Name:  "session",
			Value: uuid.NewV1().String(),
		}
		// fmt.Println(r.FormValue("eMail"), r.FormValue("password"), r.Method, "***************")
		email := r.FormValue("eMail")
		password := r.FormValue("password")
		cookie, err := h.Service.LogInUser(email, password)
		// fmt.Println(r.Method, "logPage", cookie, err)
		if err != nil {
			fmt.Println("*************************************************")
			log.Println(err)
			w.WriteHeader(400)
			user.Err = true
			user.MSG = "Email or Password is not correct"
			if err := templ.Execute(w, user); err != nil {
				w.WriteHeader(500)
				log.Println(err)
			}
			return
		} else if cookie == nil {
			fmt.Println("###################################################")
			user.Err = true
			user.MSG = "Email or Password is not correct"
			if err := templ.Execute(w, user); err != nil {
				w.WriteHeader(500)
				log.Println(err)
			}
			return
			// http.Redirect(w, r, "http://localhost:8080/logPage", 302)
			// return
		}
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "http://localhost:8080", 302)

		// if err = templ.Execute(w, nil); err != nil {
		// 	w.WriteHeader(500)
		// 	log.Println(err)
		// 	return
		// }
	default:
		w.WriteHeader(405)
		return
	}
}
