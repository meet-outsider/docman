package kit

import (
	"docman/cfg"
	"docman/pkg/log"
	"docman/pkg/model"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/mapstructure"
	"time"
)

var secret = cfg.Config.Jwt.Secret

const CLAIM = "userInfo"
const EXP = "EXP"

func GenToken(subject model.UserInfo) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		CLAIM: subject,
		EXP:   time.Now().Add(time.Hour * 24 * 3).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Error("Error generating token:", err.Error())
	}
	return tokenString
}

func ParseToken(tokenString string) (sub model.UserInfo, exp time.Time, err error) {
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
		log.Error("Error verifying token:", err.Error())
		return
	}
	// Access the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		m := claims[CLAIM].(map[string]interface{})
		// m转换成struct
		err = mapstructure.Decode(m, &sub)
		exp = time.Unix(int64(claims[EXP].(float64)), 0)

	} else {
		err = fmt.Errorf("invalid token")
	}
	return
}
