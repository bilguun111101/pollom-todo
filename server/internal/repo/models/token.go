package models

import (
	"net/http"
	"pollom/internal/shared"
)

const (
	TokenSize          = 64
	MaxTokenExpireTime = 1000 * 60 * 60 * 48

	TokenTypeVerifyEmail = "verify_email"

	TokenTypeOAuth = "oauth"
)

type Token struct {
	Token    string
	CreateAt int64
	Type     string
	Extra    string
}

func NewToken(tokentype string) *Token {
	return &Token{
		Token:    shared.NewRandomString(TokenSize),
		Type:     tokentype,
		CreateAt: shared.GetMillis(),
	}
}

func (t *Token) IsValid() *shared.AppError {
	if len(t.Token) != TokenSize {
		return shared.NewAppError("Token.IsValid", "model.token.is_valid.size", nil, "", http.StatusInternalServerError)
	}

	if t.CreateAt == 0 {
		return shared.NewAppError("Token.IsValid", "model.token.is_valid.expiry", nil, "", http.StatusInternalServerError)
	}

	return nil
}

func (t *Token) IsExpired() bool {
	if t == nil {
		return true
	}

	var expireTime int64 = MaxTokenExpireTime
	return shared.GetMillis() > (t.CreateAt + expireTime)
}
