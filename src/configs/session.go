package configs

import (
	"os"

	"github.com/gorilla/sessions"
)

func SessionInit() *sessions.CookieStore {

	store := sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SECRET")))

	return store
}
