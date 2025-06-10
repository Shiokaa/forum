package configs

import (
	"os"

	"github.com/gorilla/sessions"
)

func SessionInit() *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SECRET")))
}
