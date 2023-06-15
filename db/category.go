package db

import (
	"database/sql"
	"fmt"

	"github.com/JoseGaldamez/gambitBackend/models"
	_ "github.com/go-sql-driver/mysql"
)

// InsertCategory inserts a new category into the database
func InsertCategory(category models.Category) (int64, error) {

	fmt.Println("Begin in InsertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentence := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + category.CategoryName + "', '" + category.CategoryPath + "')"
	fmt.Println(sentence)

	var result sql.Result
	result, err = Db.Exec(sentence)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return 0, err
	}

	fmt.Println("End in InsertCategory")

	return lastInsertID, nil
}
