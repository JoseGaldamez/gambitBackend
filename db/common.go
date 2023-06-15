package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/JoseGaldamez/gambitBackend/models"
	secretmanager "github.com/JoseGaldamez/gambitBackend/secretManager"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretmanager.GetSecret(os.Getenv("SecretName"))

	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return err
	}

	fmt.Println("Connected to database")

	return nil
}

func ConnStr(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndPoint, dbName string
	dbUser = keys.Username
	authToken = keys.Password
	dbEndPoint = keys.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=True", dbUser, authToken, dbEndPoint, dbName)

	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {
	defer Db.Close()

	fmt.Println("UserIsAdmin begin")

	err := DbConnect()
	if err != nil {
		return false, "Error connecting to database: " + err.Error()
	}

	sentence := "SELECT 1 FROM users WHERE User_UUID = '" + userUUID + "' and User_Status = 1;"
	fmt.Println("Sending sentence to DB: " + sentence)

	rows, err := Db.Query(sentence)
	if err != nil {
		return false, "Error querying database: " + err.Error()
	}

	if !rows.Next() {
		return false, "User not found"
	}

	var value string
	err = rows.Scan(&value)
	fmt.Println("UserIsAdmin end > " + value)

	if err != nil {
		return false, "Error scanning rows: " + err.Error()
	}

	if value != "1" {
		return false, "User is not admin"
	}

	return true, "User is admin"
}
