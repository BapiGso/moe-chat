package mymiddleware

import "github.com/gorilla/sessions"

func test() {
	sessions.NewCookie("", "", &sessions.Options{
		Path:     "",
		Domain:   "",
		MaxAge:   0,
		Secure:   false,
		HttpOnly: false,
		SameSite: 0,
	})
}
