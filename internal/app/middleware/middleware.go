package middleware

import (
	"net/http"

	"github.com/VladimirVereshchagin/scheduler/internal/auth"
	"github.com/VladimirVereshchagin/scheduler/internal/config"
)

// Auth - проверка аутентификации
func Auth(next http.HandlerFunc, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pass := cfg.Password
		if pass == "" {
			// Пароль не задан, пропускаем
			next(w, r)
			return
		}

		// Токен из куки
		cookie, err := r.Cookie("token")
		if err != nil {
			// Нет токена, возвращаем 401
			http.Error(w, "Требуется аутентификация", http.StatusUnauthorized)
			return
		}

		// Проверяем токен
		_, err = auth.ParseToken(cookie.Value, pass)
		if err != nil {
			// Токен недействителен
			http.Error(w, "Требуется аутентификация", http.StatusUnauthorized)
			return
		}

		// Токен валиден, продолжаем
		next(w, r)
	}
}
