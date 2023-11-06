package util

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJWTFromKeyBytes(keybyte []byte) *JWT {
	// 解码给定的 PEM 数据，将其转换为一个 *pem.Block 结构
	pemblock, _ := pem.Decode(keybyte)
	if pemblock == nil {
		fmt.Println("failed to decode PEM block")
		return nil
	}
	key_private, _ := x509.ParsePKCS1PrivateKey(pemblock.Bytes)

	return NewJWT(key_private, nil)
}
func NewJWT(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *JWT {
	if publicKey == nil {
		publicKey = &privateKey.PublicKey
	}
	return &JWT{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (j *JWT) Sign(ctx context.Context, claim jwt.RegisteredClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	tokenString, _ := token.SignedString(j.privateKey)
	return tokenString
}

func (j *JWT) Verify(ctx context.Context, tokenString string) (jwt.RegisteredClaims, error) {
	var claims jwt.RegisteredClaims
	_, err := jwt.ParseWithClaims(
		tokenString,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return j.publicKey, nil
		},
		jwt.WithValidMethods([]string{"RS256"}),
	)
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}
	return claims, nil
}
