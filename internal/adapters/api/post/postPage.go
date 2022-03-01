package post

import (
	"fmt"
	entities "forum/internal/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (h *PostHandler) postPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostPage")
	cookie := r.Context().Value("user")

	templ, err := template.ParseFiles("./view/templates/postPage.html")

	post := &entities.Post{}
	commAct := &entities.Comm{}
	indPost, err := strconv.Atoi(r.RequestURI[10:])
	if err != nil {
		log.Println(err)
	}
	// commActivityInCommLikes := ""
	idComm := 0
	// fmt.Println(r.RequestURI[1:], indPost, err, "-------------------------------Post")
	if err != nil {
		fmt.Println(err)
	}
	post.ID = r.RequestURI[10:]

	post, err = h.Service.GetPost(post, indPost)

	if r.Method == "POST" {
		fmt.Println(-1, cookie)
		if cookie == nil {
			fmt.Println(2)
			http.Redirect(w, r, "http://localhost:8080/logPage", 302)
			return
		}
		userComm := r.FormValue("PostComm")
		textPostComm := r.FormValue("UserComm")
		userLike := r.FormValue("PostLike")
		userDislike := r.FormValue("PostDislike")
		commLike := r.FormValue("CommLike")
		commDislike := r.FormValue("CommDislike")
		UserId, err := h.Service.UserId(cookie.(string))
		if err != nil {
			fmt.Println(-2)
			http.Redirect(w, r, "http://localhost:8080/logPage", 302)
			log.Println(err)
			return
		}

		disOrLike := CheckLikes(userLike, userDislike)
		commActivity, idComm := CheckIdLikes(commLike, commDislike)

		//Добавляю лайки или дизлайки в таблицу ДисЛайкс
		if disOrLike == "like" && h.Service.Activity(UserId, indPost, disOrLike) == 1 {
			if err := h.Service.ReplaceAct(UserId, indPost, disOrLike, 1); err != nil {
				fmt.Println(0)
				log.Println(err)
				return
			}
			fmt.Println(1)
		} else if disOrLike == "like" && h.Service.Activity(UserId, indPost, disOrLike) == 0 {
			if err := h.Service.ReplaceAct(UserId, indPost, disOrLike, 0); err != nil {
				fmt.Println(1)
				log.Println(err)
				return
			}
		} else if disOrLike == "dislike" && h.Service.Activity(UserId, indPost, disOrLike) == 1 {
			if err := h.Service.ReplaceAct(UserId, indPost, disOrLike, 1); err != nil {
				fmt.Println(2)
				log.Println(err)
				return
			}
		} else if disOrLike == "dislike" && h.Service.Activity(UserId, indPost, disOrLike) == 0 {
			if err := h.Service.ReplaceAct(UserId, indPost, disOrLike, 0); err != nil {
				fmt.Println(3)
				log.Println(err)
				return
			}
		}

		if commActivity == "like" && h.Service.CommAct(UserId, idComm, commActivity) == 1 {
			if err := h.Service.RemoveAct(UserId, idComm, commActivity, 1); err != nil {
				fmt.Println(4)
				log.Println(err)
				return
			}
		} else if commActivity == "like" && h.Service.CommAct(UserId, idComm, commActivity) == 0 {
			if err := h.Service.RemoveAct(UserId, idComm, commActivity, 0); err != nil {
				fmt.Println(5)
				log.Println(err)
				return
			}
		} else if commActivity == "dislike" && h.Service.CommAct(UserId, idComm, commActivity) == 1 {
			if err := h.Service.RemoveAct(UserId, idComm, commActivity, 1); err != nil {
				fmt.Println(6)
				log.Println(err)
				return
			}
		} else if commActivity == "dislike" && h.Service.CommAct(UserId, idComm, commActivity) == 0 {
			if err := h.Service.RemoveAct(UserId, idComm, commActivity, 0); err != nil {
				fmt.Println(7)
				log.Println(err)
				return
			}
		} else if len(textPostComm) != 0 && userComm == "postComm" {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~COMMENT")
			if err := h.Service.AddPostComm(textPostComm, UserId, indPost); err != nil {
				log.Println(err)
				return
			}
		}

	}
	// Считаем лайки и дизлайки Поста
	post, err = h.Service.GetLikes(indPost, post)
	post, err = h.Service.GetDislikes(indPost, post)

	// Считаем лайки и дизлайки Комментов
	commAct, err = h.Service.GetCommLikes(idComm, commAct)
	commAct, err = h.Service.GetCommDislikes(idComm, commAct)

	post, err = h.Service.GetAllComm(indPost, post, commAct)

	fmt.Println(post.Title, post.Body, commAct.Dislikes, commAct.Likes, "###############################")
	templ.Execute(w, post)
}
