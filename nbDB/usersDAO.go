package nbDB

import (
	"fmt"
	"time"

	"github.com/swathiGiligar/nbServer/resources"
)

type DbUser struct {
	DbUserId      int64
	DbUserName    string
	DbLoginId     string
	DbAccessLevel int
	DbLastLogin   time.Time
}

func InsertNewUser(newUser DbUser) {
	db := connectToDB()

	// close database
	defer db.Close()

	result, e := db.Exec(resources.InsertNewUserStmt, newUser.DbUserName,
		newUser.DbLoginId, newUser.DbAccessLevel)

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\nRows Inserted = %d", rowsAffected)

	CheckError(e)
}

func FetchUser(loginId string) DbUser {
	// open database
	db := connectToDB()

	// close database
	defer db.Close()

	newUser := DbUser{}

	err := db.QueryRow(resources.FetchUserWithLoginIdStmt, loginId).Scan(
		&newUser.DbUserId, &newUser.DbUserName, &newUser.DbLoginId,
		&newUser.DbAccessLevel, &newUser.DbLastLogin)
	CheckError(err)

	fmt.Println(newUser.DbUserId, newUser.DbUserName, newUser.DbLoginId,
		newUser.DbAccessLevel, newUser.DbLastLogin)

	CheckError(err)
	return newUser

}

func UpdateLastLogin(userId int64) {
	// open database
	db := connectToDB()

	// close database
	defer db.Close()

	result, e := db.Exec(resources.UpdateLastLoginStmt, userId)
	CheckError(e)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\nRows Affected = %d", rowsAffected)
	fmt.Printf("\nRows Affected = %d", rowsAffected)

}

func UpdateUserAccess(userId int64, accessLevel int) {
	// open database
	db := connectToDB()

	// close database
	defer db.Close()

	result, e := db.Exec(resources.UpdateUserAccessStmt, accessLevel, userId)
	CheckError(e)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\nRows Affected = %d", rowsAffected)

}
