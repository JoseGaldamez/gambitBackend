package routers

import (
	"encoding/json"
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

	isAdmin, msg := db.UserIsAdmin(user)
	if !isAdmin {
		return 400, msg
	}

	result, errInsert := db.InsertCategory(category)
	if errInsert != nil {
		return 400, "Something went wrong inserting: " + body + errInsert.Error()
	}

	return 200, "{CategoryID: " + strconv.Itoa(int(result)) + "}"
}
