package nbDB

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/swathiGiligar/nbServer/resources"
)

type DbNotice struct {
	DbNoticeId   int64
	DbHeading    string
	DbPrice      string
	DbCategory   string
	DbAreaLevel1 string
	DbAreaLavel2 string
	DbContact    string
	DbDetails    string
	DbCreatedOn  time.Time
	DbUpdatedOn  time.Time
}

func UpdateNotice(updatedNotice DbNotice) {
	db := connectToDB()

	// close database
	defer db.Close()

	result, e := db.Exec(resources.UpdateNoticeStmt, updatedNotice.DbHeading,
		updatedNotice.DbPrice, updatedNotice.DbCategory,
		updatedNotice.DbAreaLevel1, updatedNotice.DbAreaLavel2,
		updatedNotice.DbContact, updatedNotice.DbDetails, time.Now(),
		updatedNotice.DbNoticeId)
	CheckError(e)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\nRows Updated = %d", rowsAffected)
}

func InsertNotice(newNotice DbNotice) {
	db := connectToDB()

	// close database
	defer db.Close()

	result, e := db.Exec(resources.InsertNewNoticeStmnt, newNotice.DbHeading, newNotice.DbPrice,

		newNotice.DbCategory, newNotice.DbAreaLevel1, newNotice.DbAreaLavel2,
		newNotice.DbContact, newNotice.DbDetails, "ACTIVE")

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\nRows Inserted = %d", rowsAffected)

	CheckError(e)
}

func CloseNotice(noticeId int64) {
	// open database
	db := connectToDB()

	// close database
	defer db.Close()

	result, e := db.Exec(resources.CloseNoticeStmt, time.Now(), noticeId)
	CheckError(e)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\nRows Affected = %d", rowsAffected)

}

func FetchNotices(status string) []DbNotice {
	// open database
	db := connectToDB()

	// close database
	defer db.Close()

	rows, err := db.Query(resources.FetchNoticesStmt, "ACTIVE")
	CheckError(err)

	defer rows.Close()
	var notices = []DbNotice{}
	for rows.Next() {
		var noticeId int64
		var heading string
		var price string
		var category string
		var area_level_1 string
		var area_level_2 string
		var contact string
		var details string
		var created_on time.Time

		err = rows.Scan(&noticeId, &heading, &price, &category, &area_level_1,
			&area_level_2, &contact, &details, &created_on)
		CheckError(err)
		currentNotice := DbNotice{DbNoticeId: noticeId, DbHeading: heading,
			DbPrice:    price,
			DbCategory: category, DbAreaLevel1: area_level_1,
			DbAreaLavel2: area_level_2, DbContact: contact, DbDetails: details, DbCreatedOn: created_on}
		notices = append(notices, currentNotice)

		fmt.Println(noticeId, heading, price, category, area_level_1,
			area_level_2, contact, details)
	}

	CheckError(err)
	return notices

}

func connectToDB() *sql.DB {
	psqlconn :=
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			resources.Host, resources.Port, resources.User, resources.Password,
			resources.Dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
