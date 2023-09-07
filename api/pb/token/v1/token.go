package token

import (
	"net/http"
	"net/url"
)

func (c *Token) SetCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "ekube.access_token",
		Value:    url.QueryEscape(c.AccessToken),
		MaxAge:   0,
		Path:     "/",
		Domain:   "",
		SameSite: http.SameSiteDefaultMode,
		Secure:   false,
		HttpOnly: true,
	})
}
