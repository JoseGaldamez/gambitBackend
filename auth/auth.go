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
	fmt.Println("Part to decode: " + parts[1])

	userInfo, err := base64.RawStdEncoding.DecodeString(parts[1] + "/fg==")
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

// .eyJzdWIiOiIzN2YwNzkyNS0zNmIyLTRkOTYtOGMzMi04MDFmYTk4MGJhZjIiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV9CSVloU2NBbFgiLCJ2ZXJzaW9uIjoyLCJjbGllbnRfaWQiOiIzOWFibWxmNmplcWZvZWR0ZHI2NDVudWF0IiwiZXZlbnRfaWQiOiJkMGNhNDRjZC05OWNkLTRhNjQtYmEyNy1lZGYzZDk0MWRkMDUiLCJ0b2tlbl91c2UiOiJhY2Nlc3MiLCJzY29wZSI6ImF3cy5jb2duaXRvLnNpZ25pbi51c2VyLmFkbWluIG9wZW5pZCBlbWFpbCIsImF1dGhfdGltZSI6MTY4Njg2MTUyNiwiZXhwIjoxNjg2OTQ3OTI2LCJpYXQiOjE2ODY4NjE1MjYsImp0aSI6IjJlODA5NTg1LWQyOGEtNDA1OS1iZDI1LTU5NDc3Y2UzZDc1NiIsInVzZXJuYW1lIjoiMzdmMDc5MjUtMzZiMi00ZDk2LThjMzItODAxZmE5ODBiYWYyIn0.
