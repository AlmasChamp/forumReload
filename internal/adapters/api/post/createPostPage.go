package post

import (
	"fmt"
	"html/template"
	"net/http"
)

func (h *PostHandler) createPostPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreatePostPage")
	templ, err := template.ParseFiles("./view/templates/createPostPage.html")
	if err != nil {
		fmt.Println(err)
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

	templ.Execute(w, nil)
}
