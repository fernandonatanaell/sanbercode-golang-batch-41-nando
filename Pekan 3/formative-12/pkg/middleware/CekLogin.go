package middleware

import (
	"net/http"
)

// Middleware untuk check login
func CekLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			uname, pass, ok := r.BasicAuth()
			if !ok || (uname == "" || pass == "") {
				w.Write([]byte("Username atau Password tidak boleh kosong"))
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
