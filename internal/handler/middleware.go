package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gitkoDev/KODE-test-task/helpers"
)


func (h *Handler) IdentifyUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Values("Authorization")

		if authToken == nil {
			helpers.RespondWithError(w, errors.New("empty authorization header"), http.StatusUnauthorized)
			return
		}

		authToken = strings.Split(authToken[0], " ")
		tokenPart := authToken[1]

		user_id, err := h.service.Auth.ParseToken(tokenPart)
		if err != nil {
			helpers.RespondWithError(w, err, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", user_id)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
