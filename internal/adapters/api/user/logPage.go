package user

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func (h *Handler) logPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logPage", r.Method)
	templ, err := template.ParseFiles("./view/templates/loginPage.html")
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	switch r.Method {
	case "GET":
		if r.URL.Path != "/logPage" {
			w.WriteHeader(404)
			return
		}
		// er := entities.Error{}
		if err = templ.Execute(w, nil); err != nil {
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
		fmt.Println(r.FormValue("eMail"), r.FormValue("password"), r.Method, "***************")
		email := r.FormValue("eMail")
		password := r.FormValue("password")
		cookie, err := h.Service.LogInUser(email, password)
		// fmt.Println(r.Method, "logPage", cookie, err)
		if err != nil || cookie == nil {
			fmt.Println("Check")
			log.Println(err, cookie)
			w.WriteHeader(404)
			return
		}
		http.SetCookie(w, cookie)

		http.Redirect(w, r, "http://localhost:8080", 302)

		if err = templ.Execute(w, nil); err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}
	default:
		w.WriteHeader(405)
		return
	}
}
