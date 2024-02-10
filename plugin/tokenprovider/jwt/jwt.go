package jwt

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"go-200lab-g09/common"
	"go-200lab-g09/plugin/tokenprovider"

	"github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	name   string
	secret string
}

func NewJWTProvider(name string) *jwtProvider {
	return &jwtProvider{name: name}
}

func (p *jwtProvider) GetPrefix() string {
	return p.Name()
}

func (p *jwtProvider) Get() interface{} {
	return p
}

func (p *jwtProvider) Name() string {
	return p.name
}

func (p *jwtProvider) InitFlags() {
	flag.StringVar(&p.secret, "jwt-secret", common.GenSalt(40), "Secret key for generating JWT")
}

func (p *jwtProvider) Configure() error {
	return nil
}

func (p *jwtProvider) Run() error {
	return nil
}

func (p *jwtProvider) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		c <- true
	}()

	return c
}

type myClaims struct {
	Payload common.TokenPayLoad `json:"payload"`
	jwt.StandardClaims
}

type token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

func (t *token) GetToken() string {
	return t.Token
}

func (j *jwtProvider) SecretKey() string {
	return j.secret
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (tokenprovider.Token, error) {
	now := time.Now()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		common.TokenPayLoad{
			UId:   data.UserId(),
			URole: data.Role(),
		},
		jwt.StandardClaims{
			ExpiresAt: now.Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  now.Local().Unix(),
			Id:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &token{
		Token:   myToken,
		Expiry:  expiry,
		Created: now,
	}, nil
}

func (j *jwtProvider) Validate(mytoken string) (tokenprovider.TokenPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, tokenprovider.ErrInvalidToken
		}
		return []byte(j.secret), nil
	}

	jwtToken, err := jwt.ParseWithClaims(mytoken, &myClaims{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, tokenprovider.ErrExpiredToken) {
			return nil, tokenprovider.ErrExpiredToken
		}

		return nil, tokenprovider.ErrInvalidToken
	}

	if !jwtToken.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := jwtToken.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	return claims.Payload, nil
}
