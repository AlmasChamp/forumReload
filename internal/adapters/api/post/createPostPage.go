package post

import (
	"fmt"
	entities "forum/internal/model"
	"html/template"
	"log"
	"net/http"
)

func (h *PostHandler) createPostPage(w http.ResponseWriter, r *http.Request) {
	cookie := r.Context().Value("user")

	templ, err := template.ParseFiles("./view/templates/createPostPage.html")
	if err != nil {
		fmt.Println("1")
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	categories := []string{"Тренировки", "Соревнования", "Отдых"}

	user := entities.User{
		Categories: categories,
	}
	switch r.Method {
	case "GET":
		if cookie == nil {
			fmt.Println("2")
			http.Redirect(w, r, "http://localhost:8080/logPage", 302)
		}
		if r.URL.Path != "/createPostPage" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			log.Println(err)
			return
		}
		if err := templ.Execute(w, user); err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			log.Println(err)
			return
		}
	case "POST":
		fmt.Println("3")
		title := r.FormValue("title")
		body := r.FormValue("body")
		if r.URL.Path != "/createPostPage" {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			log.Println(err)
			return
		} else if cookie == nil {
			http.Redirect(w, r, "http://localhost:8080/", 302)
		} else if cookie != nil {
			fmt.Println("4")
			cookie, err := h.Service.ValueCookie(cookie.(string))
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println("5")
			userId, err := h.Service.UserId(cookie)
			if err != nil {
				log.Println(err)
				return
			}

			if err := h.Service.CreatePost(title, body, userId); err != nil {
				log.Println(err)
				return
			}

			fmt.Println("Redirect")
			http.Redirect(w, r, "http://localhost:8080", 302)
		}

	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

}

// if r.Method == "POST" {
// 	userCookie, err := r.Cookie("session")
// 	if err != nil {
// 		fmt.Println("Error from Create Post")
// 	}
// 	d.DB.Exec(`INSERT INTO Posts (title,body,userId)
// 	VALUES($1,$2,$3)`, r.FormValue("title"), r.FormValue("body"), getUserIdFromCookie(d, userCookie.Value))
// 	http.Redirect(w, r, "http://localhost:8080", 302)
// }

// templ.Execute(w, user)

// else if cookieFromDb == "" || cookieFromDb != cookie {
// 	fmt.Println(cookieFromDb, "CREATEPOSTERROR")
// 	http.Redirect(w, r, "http://localhost:8080/createPostPage", 302)
// }

// if err := templ.Execute(w, user); err != nil {
// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
// 	log.Println(err)
// 	return
// }
// }
// 	userInfo.Islogged = true
// } else if cookieFromDb != cookie {
// 	http.Redirect(w, r, "http://localhost:8080/", 302)
// }

// d.DB.Exec(`INSERT INTO Posts (title,body,userId)
// VALUES($1,$2,$3)`, r.FormValue("title"), r.FormValue("body"), getUserIdFromCookie(d, userCookie.Value))
// http.Redirect(w, r, "http://localhost:8080", 302)
// if err := templ.Execute(w, user); err != nil {
// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
// 	log.Println(err)
// 	return
