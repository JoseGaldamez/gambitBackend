package tools

import (
	"fmt"
	"os"
	"time"
)

func DateMySQL() string {
	currentDate := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", currentDate.Year(), currentDate.Month(), currentDate.Day(), currentDate.Hour(), currentDate.Minute(), currentDate.Second())
}

func ValidateEnvironment() bool {
	_, ok := os.LookupEnv("SecretName")
	if !ok {
		return ok
	}

	// _, ok = os.LookupEnv("UserPoolId")
	// if !ok {
	// 	return ok
	// }

	// _, ok = os.LookupEnv("Region")
	// if !ok {
	// 	return ok
	// }

	_, ok = os.LookupEnv("UrlPrefix")
	if !ok {
		return ok
	}

	return ok
}
