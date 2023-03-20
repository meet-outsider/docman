package jwt

import (
	conf "docman/config"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var secret = conf.Config.Jwt.Secret

const CLAIM = "UserId"
const EXP = "EXP"

func GenToken(userId int) string {
	//secret := conf.Config.Jwt.Secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		CLAIM: userId,
		EXP:   time.Now().Add(time.Hour * 24 * 3).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Error creating token:", err)
	}
	return tokenString
}

func ParseToken(tokenString string) (userId int, exp time.Time, err error) {
	// Verify the token with the same signing key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used to sign the token
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("Error verifying token:", err)
		return
	}
	// Access the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdFl, _ := claims[CLAIM].(float64)
		userId = int(userIdFl)
		exp = time.Unix(int64(claims[EXP].(float64)), 0)
	} else {
		err = fmt.Errorf("invalid token")
	}
	return
}
