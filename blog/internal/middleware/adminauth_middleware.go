package middleware

import (
	"blog/internal/models"
	"fmt"
	"net/http"
)

type AdminAuthMiddleware struct {
}

func NewAdminAuthMiddleware() *AdminAuthMiddleware {
	return &AdminAuthMiddleware{}
}

func (m *AdminAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		id := r.Header.Get("identity")
		user, ok := models.GetById(id)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("用户不存在!"))
			return
		}
		if user.Role != 9999 && user.Role != 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("您没有权限!"))
			return
		}
		fmt.Println(id, user.Role)

		// Passthrough to next handler if need
		next(w, r)
	}
}
