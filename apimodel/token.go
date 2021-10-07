package apimodel

import (
	"time"

	"github.com/go-openapi/strfmt"
)

func NewToken(domain string, redirectUrl string, clientSecret string, clientID string, accessToken string, acceessTokenExp int64, refreshToken string) *Token {
	return &Token{
		Domain:            domain,
		RedirectURI:       redirectUrl,
		ClientSecret:      clientSecret,
		ClientID:          clientID,
		AccessToken:       accessToken,
		AccessTokenExpire: strfmt.DateTime(time.Now().Add(time.Duration(acceessTokenExp) * time.Second)),
		RefreshToken:      refreshToken,
	}
}
