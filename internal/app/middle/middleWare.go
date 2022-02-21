package middle

import (
	"context"
	"fmt"
	"net/http"
)

func MiddleWare(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MiddleWare")
		cookie, err := r.Cookie("session")
		if err != nil {
			fmt.Println("MiddleWareERROR")
			// handler.ServeHTTP(w, r)
			handler(w, r)
			return
		} else {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "user", cookie.Value)
			r = r.WithContext(ctx)
			// handler.ServeHTTP(w, r)
			handler(w, r)
		}
	}
}
