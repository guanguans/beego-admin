package helpers

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
)

// https://gist.github.com/cryptix/45c33ecf0ae54828e63b
const (
	privateKeyPath = "storage/jwt/jwt-key.pem"     // openssl genrsa -out jwt-key.pem keysize(1024)
	publicKeyPath  = "storage/jwt/jwt-key.pem.pub" // openssl rsa -in jwt-key.pem -pubout > jwt-key.pem.pub
)

var (
	verifyKey    *rsa.PublicKey
	mySigningKey *rsa.PrivateKey
)

type EasyToken struct {
	Username string
	Expires  int64
}

func init() {
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}

	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	mySigningKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}
}

func (e EasyToken) GenerateToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: e.Expires,
		Issuer:    e.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, err
}

func (e EasyToken) ParseToken(tokenString string) (string, error) {
	if tokenString == "" {
		return "", errors.New("Token cannot be empty!")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if token == nil {
		return "", errors.New("not work")
	}
	if !token.Valid {
		return "", err
	}

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", errors.New("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return "", errors.New("Timing is everything")
		} else {
			return "", err
		}
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["iss"].(string), nil
}
