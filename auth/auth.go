package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJson struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidateToken(token string) (bool, string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, "Invalid token structure", nil
	}

	fmt.Println(parts)
	fmt.Println(parts[1])

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Error decoding token", err.Error())
		return false, "Decoding error", err
	}

	var tkj TokenJson
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("Error unmarshalling token", err.Error())
		return false, "Unmarshalling error", err
	}

	rightNowTime := time.Now()
	tokenTimeExpire := time.Unix(int64(tkj.Exp), 0)

	if tokenTimeExpire.Before(rightNowTime) {
		fmt.Println("Token expired: " + tokenTimeExpire.String())
		return false, "Token expired", nil
	}

	return true, string(tkj.Username), nil

}
