package user

import (
	"database/sql"
	"errors"
	"fmt"
	entities "forum/internal/model"
	"log"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetValueCookie(userCookie string) (string, error) {
	cookieFromDb := ""
	err := r.db.QueryRow(`SELECT value
	FROM Cookie
	WHERE value = ?`, userCookie).Scan(&cookieFromDb)
	if err != nil {
		return "", err
	}
	return cookieFromDb, nil
}

func (r *Repository) GetAllPosts() ([]entities.Post, error) {
	outInfo := entities.Out{}
	rows, err := r.db.Query(`SELECT id,title,body
	FROM Posts`)

	if err != nil {
		// fmt.Println("Tuta")
		log.Println(err)
		return outInfo.UsersPosts, err
	}

	defer rows.Close()

	for rows.Next() {
		userPost := entities.Post{}
		err := rows.Scan(&userPost.ID, &userPost.Title, &userPost.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outInfo.UsersPosts = append(outInfo.UsersPosts, userPost)
	}
	return outInfo.UsersPosts, nil
}

func (r *Repository) AddUser(name, email, password string) error {
	_, err := r.db.Exec(`INSERT INTO Users (userName,eMail,password)
		VALUES ($1,$2,$3)`, name, email, password)
	if err != nil {
		fmt.Println("InsertUserERROR")
		return errors.New("Sorry name or login is already taken")
	}
	fmt.Println("InsertUser")
	return nil
}

func (r *Repository) GetUserPassword(email string) (string, error) {
	password := ""
	rows, err := r.db.Query(`SELECT password
		FROM Users
		WHERE eMail=?`, email)
	if err != nil {
		log.Println(err)
		return "", err
	}
	for rows.Next() {
		rows.Scan(&password)
	}
	return password, nil
}

func (r *Repository) GetUserId(email string) (string, error) {
	userId := ""
	err := r.db.QueryRow(`SELECT id
	FROM Users
	WHERE eMail = ?`, email).Scan(&userId)
	if err != nil {
		return "", err
	}
	// fmt.Println(userId, "UseridFromUsers")
	return userId, nil
}

func (r *Repository) SetCookie(cookieVal string, cookieExp int, id string) error {
	_, err := r.db.Exec(`INSERT OR REPLACE INTO Cookie (value,expires,userId)
	VALUES ($1,$2,$3)`, cookieVal, cookieExp, id)
	if err != nil {
		return err
	}
	// _, err := r.db.Exec(`INSERT OR REPLACE INTO Cookie (value,expires,userId)
	// VALUES ($1,$2,$3)`, cookieVal, cookieExp, id)
	// if err != nil {
	// 	return err
	// }
	// d.DB.Exec(`INSERT OR REPLACE INTO DisLikes (id,like,userId,postIdInComm)
	// VALUES ((SELECT id
	// 	FROM DisLikes
	// 	WHERE userId = $1 AND postIdInComm = $2),$3,$4,$5)`, UserId, indPost, nil, UserId, indPost)

	// _, err = r.db.Exec("INSERT INTO Cookie (Value, Expires,UserId)VALUES(?,?,?)", cookie.Value, cookie.Expires, id)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	return nil
}

func (r *Repository) DeleteCookie(cookieVal string) error {
	_, err := r.db.Exec(`DELETE FROM Cookie
	WHERE value = ?`, cookieVal)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
