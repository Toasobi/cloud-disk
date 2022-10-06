package middleware

import (
	"cloud-disk/user-api/helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorization"))
			return
		}

		userclaim, err := helper.AnalyseToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		r.Header.Set("UserId", string(rune(userclaim.Id)))
		r.Header.Set("UserIdentity", userclaim.Identity)
		r.Header.Set("UserName", userclaim.Name)

		next(w, r)
	}
}
