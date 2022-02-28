package main

import (
	"context"
	"fmt"
	"forum/internal/adapters/repository"
	"forum/internal/app"
	"forum/internal/domain/user"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Init DataBase
	db, err := repository.InitDb()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go user.DeleteCookie(db, ctx)
	// CreateTables
	if err := repository.CreateTables(db); err != nil {
		log.Fatal(err)
		return
	}

	mux := http.NewServeMux()


	userComposite, err := app.NewUserComposites(db)
	postComposite, err := app.NewPostComposites(db)

	userComposite.Handler.Register(mux)
	postComposite.Handler.Register(mux)

	fmt.Println("Start on port 8080")
	http.ListenAndServe(":8080", mux)
	cancel()
	fmt.Println("Priehali")

	// Init Storage
	// repository := userRep.NewRepository(db)
	// // Init Service
	// service := user.NewService(repository)
	// // Init Handler
	// handler := userApi.NewHandler(service)
	// Init Mux

}
