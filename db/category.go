package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/JoseGaldamez/gambitBackend/models"
	"github.com/JoseGaldamez/gambitBackend/tools"
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

func UpdateCategory(category models.Category) error {

	fmt.Println("Begin in UpdateCategory")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentence := "UPDATE category SET "

	if len(category.CategoryName) > 0 {
		sentence += " Categ_Name = '" + tools.StringScape(category.CategoryName) + "'"
	}

	if len(category.CategoryPath) > 0 {
		if !strings.HasSuffix(sentence, "SET ") {
			sentence += ", "
		}
		sentence += "Categ_Path = '" + tools.StringScape(category.CategoryPath) + "'"

	}

	sentence += " WHERE Categ_ID = " + strconv.Itoa(category.CategoryID)

	fmt.Println(sentence)

	_, err = Db.Exec(sentence)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return err
	}

	fmt.Println("End in UpdateCategory - Successfully")

	return nil
}
