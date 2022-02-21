package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func DeleteCookie(db *sql.DB, ctx context.Context) {
loop:
	for {
		select {
		case <-time.After(60 * time.Second):
			_, err := db.Exec("DELETE FROM Cookie WHERE expires <?", time.Now())
			if err != nil {
				log.Println("DeleteCookie")
				continue
			}
			continue
		case <-ctx.Done():
			fmt.Println("Hello I am GoRoutine")
			break loop
		}
	}
}
