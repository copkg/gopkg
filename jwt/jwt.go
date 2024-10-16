package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	TokenInvalid = errors.New("token invalid")
)

type JWT struct {
	SigningKey    string
	ExpiresTime   int64
	SigningMethod string
}
type CustomClaims struct {
	StaffNo string `json:"staff_no,omitempty"`
	Name    string `json:"name,omitempty"`
	AppID   int    `json:"app_id,omitempty"`
	jwt.RegisteredClaims
}

func NewJWT(c JWT) *JWT {
	return &JWT{
		SigningKey:    c.SigningKey,
		ExpiresTime:   c.ExpiresTime,
		SigningMethod: c.SigningMethod,
	}
}

func (j *JWT) CreateClaims() CustomClaims {
	var claims = CustomClaims{}
	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	var signingMethod jwt.SigningMethod
	switch j.SigningMethod {
	case "HS256":
		signingMethod = jwt.SigningMethodHS256
	case "HS512":
		signingMethod = jwt.SigningMethodHS512
	case "HS384":
		signingMethod = jwt.SigningMethodHS384
	}
	token := jwt.NewWithClaims(signingMethod, claims)
	return token.SignedString([]byte(j.SigningKey))
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(j.SigningKey), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
