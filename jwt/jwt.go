package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	TokenInvalid = errors.New("token invalid")
)

type JWTConf struct {
	SigningKey    string
	ExpiresTime   int64
	SigningMethod string
	Issuer        string
}

var JWT *JWTConf

type CustomClaims struct {
	UID     uint   `json:"uid"`
	ID      int64  `json:"id"`
	Role    string `json:"role,omitempty"`
	StaffNo string `json:"staffNo,omitempty"`
	AppID   int    `json:"appId,omitempty"`
	UserID  string `json:"userId,omitempty"`
	Name    string `json:"name,omitempty"`
	Avatar  string `json:"avatar,omitempty"`
	Type    int8   `json:"type,omitempty"`
	jwt.RegisteredClaims
}

func NewJWT(c *JWTConf) *JWTConf {
	return &JWTConf{
		SigningKey:    c.SigningKey,
		ExpiresTime:   c.ExpiresTime,
		SigningMethod: c.SigningMethod,
		Issuer:        c.Issuer,
	}
}

func (j *JWTConf) CreateClaims(claims CustomClaims) CustomClaims {
	return claims
}

// 创建一个token
func (j *JWTConf) CreateToken(claims jwt.Claims) (string, error) {
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
func (j *JWTConf) ParseToken(tokenString string) (*CustomClaims, error) {
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
