package handlers

import (
	"fmt"
	"strconv"

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
	fmt.Println("ID: " + id)

	idn, _ := strconv.Atoi(id)
	fmt.Printf("IDN: %d\n", idn)

	return 400, "Method not found"
}
