package models

import "pollom/internal/shared"

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

func (t *Token) IsExpired() bool {
	if t == nil {
		return true
	}

	var expireTime int64 = MaxTokenExpireTime
	return shared.GetMillis() > (t.CreateAt + expireTime)
}
