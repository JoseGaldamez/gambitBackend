package routers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/JoseGaldamez/gambitBackend/db"
	"github.com/JoseGaldamez/gambitBackend/models"
)

func InsertCategory(body string, user string) (int, string) {

	var category models.Category

	err := json.Unmarshal([]byte(body), &category)
	if err != nil {
		return 400, "Invalida JSON"
	}

	if len(category.CategoryName) < 1 {
		return 400, "Category Name is empty"
	}

	if len(category.CategoryPath) < 1 {
		return 400, "Category Path is empty"
	}

	var userIsAdmin bool
	var message string

	userIsAdmin, message = db.UserIsAdmin(user)

	fmt.Println("UserIsAdmin: OK")
	fmt.Println("Message: " + message)
	fmt.Println(userIsAdmin)

	if !userIsAdmin {
		return 400, message
	}

	result, errInsert := db.InsertCategory(category)
	if errInsert != nil {
		return 400, "Something went wrong inserting: " + body + errInsert.Error()
	}

	return 200, "{CategoryID: " + strconv.Itoa(int(result)) + "}"
}

func UpdateCategory(body string, user string, id int) (int, string) {

	var category models.Category

	err := json.Unmarshal([]byte(body), &category)
	if err != nil {
		return 400, "Invalida JSON"
	}

	if len(category.CategoryName) < 1 && len(category.CategoryPath) < 1 {
		return 400, "Category Name and Category Path is empty"
	}

	var userIsAdmin bool
	var message string

	userIsAdmin, message = db.UserIsAdmin(user)

	fmt.Println("UserIsAdmin: OK")
	fmt.Println("Message: " + message)
	fmt.Println(userIsAdmin)

	if !userIsAdmin {
		return 400, message
	}

	category.CategoryID = id

	errUpdate := db.UpdateCategory(category)

	if errUpdate != nil {
		return 400, "Something went wrong updating: " + body + errUpdate.Error()
	}

	return 200, "{ updated: true }"
}

func DeleteCategory(body string, user string, id int) (int, string) {

	if id == 0 {
		return 400, "Category ID is empty"
	}

	var userIsAdmin bool
	var message string

	userIsAdmin, message = db.UserIsAdmin(user)

	fmt.Println("UserIsAdmin: OK")
	fmt.Println("Message: " + message)
	fmt.Println(userIsAdmin)

	if !userIsAdmin {
		return 400, message
	}

	err := db.DeleteCategory(id)

	if err != nil {
		return 400, "Something went wrong deleting: " + err.Error()
	}

	return 400, "{ deleted: true, id: " + strconv.Itoa(id) + ", message: 'Category deleted' }"
}
