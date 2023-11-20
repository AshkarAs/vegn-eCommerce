package jwttoken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken(secretKey string) (string, error) {

	claims := jwt.MapClaims{
		"exp": time.Now().Unix() + 3600000,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error occured while creating token:", err)
		return "", err
	}

	return signedToken, nil

}


func GenerateAcessToken(securityKey string, id string) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Unix() + 300,
		"id":  id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(securityKey))
	if err != nil {
		fmt.Println(err, "Error creating acesss token ")
		return "", err
	}
	return tokenString, nil
}



func TempTokenForOtpVerification(securityKey string, phone string) (string, error) {
	claims := jwt.MapClaims{
		"phone": phone,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(securityKey))
	if err != nil {
		fmt.Println(err, "error at creating jwt token ")
	}
	return tokenString, err
}

func UnbindPhoneFromClaim(tokenString string, tempVerificationKey string) (string, error) {
	
	secret := []byte(tempVerificationKey)
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !parsedToken.Valid {
		fmt.Println(err)
		return "", err
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	phone := claims["phone"].(string)

	return phone, nil
}

