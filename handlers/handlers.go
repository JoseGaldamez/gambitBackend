package handlers

import (
	"fmt"
	"strconv"

	"github.com/JoseGaldamez/gambitBackend/auth"
	"github.com/JoseGaldamez/gambitBackend/routers"
	"github.com/aws/aws-lambda-go/events"
)

func Handlers(
	path string,
	method string,
	body string,
	headers map[string]string,
	request events.APIGatewayV2HTTPRequest) (int, string) {

	fmt.Println("Processing request" + path + " > " + method)
	fmt.Println("Body: " + body)

	id := request.PathParameters["id"]

	idn, _ := strconv.Atoi(id)
	fmt.Printf("IDN: %d\n", idn)

	isOk, code, user := validateAuthorization(path, method, headers)
	if !isOk {
		return code, user
	}

	switch path[0:5] {
	case "user":
		return handleUsers(body, path, method, user, id, request)
	case "prod":
		return handleProducts(body, path, method, user, idn, request)
	case "stoc":
		return handleStocks(body, path, method, user, idn, request)
	case "addr":
		return handleAddresses(body, path, method, user, idn, request)
	case "/cate":
		return handleCategories(body, path, method, user, idn, request)
	case "orde":
		return handleOrdens(body, path, method, user, idn, request)

	}

	return 400, "Method not found"
}

func validateAuthorization(path string, method string, headers map[string]string) (bool, int, string) {

	fmt.Println("Validating authorization")
	fmt.Println("Path: " + path)
	fmt.Println("Method: " + method)
	fmt.Println(headers)

	if (path == "product" && method == "GET") || (path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token not found"
	}

	isOk, message, err := auth.ValidateToken(token)
	if !isOk {

		if err != nil {
			fmt.Println("Error validating token", err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Token validated")
			return false, 401, message
		}

	}

	return true, 200, message
}

func handleUsers(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method not found - handleUser"
}

func handleProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method not found - handleProducts"
}

func handleCategories(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	}
	return 400, "Method not found - handleCategories"
}

func handleStocks(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method not found - handleStocks"
}

func handleAddresses(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method not found - handleAddresses"
}

func handleOrdens(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method not found - handleOrdens"
}
