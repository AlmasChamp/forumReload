package user

import (
	"fmt"
	entities "forum/internal/model"
	"html/template"
	"log"
	"net/http"
)

// mainPage
func (h *Handler) mainPage(w http.ResponseWriter, r *http.Request) {

	cookie := r.Context().Value("user")

	if r.URL.Path != "/" {
		w.WriteHeader(404)
		return
	}

	// fmt.Println("mainPAge", r.Method)

	if r.Method != http.MethodGet {
		// deleteUserCookie := r.FormValue("logOut")
		if r.Method == "POST" && r.FormValue("logOut") == "LogOut" {
			delCookie := h.Service.LogOut(cookie.(string))
			if delCookie == nil {
				fmt.Errorf("Cookie doesn't delete")
				return
			}
			http.SetCookie(w, delCookie)
			http.Redirect(w, r, "http://localhost:8080/", 302)
		}
		w.WriteHeader(405)
		return
	}

	temp, err := template.ParseFiles("./view/templates/mainPage.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	userInfo := entities.Out{}
	userInfo.UsersPosts, err = h.Service.AllPost()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	if cookie == nil {
		fmt.Println("Here")
		temp.Execute(w, userInfo)
		return
	} else if cookie != nil {
		cookieFromDb, err := h.Service.ValueCookie(cookie.(string))
		if err != nil {
			temp.Execute(w, userInfo)
			log.Println(err)
			return
		}
		fmt.Println(cookie, cookieFromDb, "********************************")

		if cookieFromDb == cookie {
			fmt.Println("alhamduliLlyah cookie from DB")
			userInfo.Islogged = true
		} else if cookieFromDb == "" || cookieFromDb != cookie {
			userInfo.Islogged = false
			http.Redirect(w, r, "http://localhost:8080/", 302)
		}
		// fmt.Println(infoAboutFromUser)
		temp.Execute(w, userInfo)

	}
}
