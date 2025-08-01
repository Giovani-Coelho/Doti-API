package middleware

import (
	"net/http"

	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/auth"
)

func EnsureAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := auth.GetAuthenticatedUser(r)

		if err != nil {
			res := resp.NewHttpJSONResponse(w)
			res.AddBody(map[string]string{
				"error":   err.Error(),
				"message": "Please sign in to access this resource",
			})
			res.Write(401)
			return
		}

		ctxWithUser := auth.SetUserInContext(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctxWithUser))
	})
}
