package nbDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "notice_board"

	FetchNoticesStmt = `SELECT "notice_id", "heading", "price", "category",
	"area_level_1", "area_level_2", "contact", "details"
	 FROM "notices" WHERE status=$1 `

	insertNewNoticeStmnt = `insert into "notices"("heading", "price",
	 "category", "area_level_1", "area_level_2", "contact", "details", 
	 "status") values($1, $2, $3, $4, $5, $6, $7, $8)`

	closeNoticeStmt = `update "notices" set "status"='CLOSED'  
	where "notice_id"=$1`

	updateNoticeStmt = `update "notices" set "heading"=$1, "price"=$2,
	"category"=$3, "area_level_1"=$4, "area_level_2"=$5, "contact"=$6, 
	"details"=$7 where "notice_id"=$8`
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
}

func UpdateNotice(updatedNotice DbNotice) {
	db := connectToDB()

	// close database
	defer db.Close()

	result, e := db.Exec(updateNoticeStmt, updatedNotice.DbHeading,
		updatedNotice.DbPrice, updatedNotice.DbCategory,
		updatedNotice.DbAreaLevel1, updatedNotice.DbAreaLavel2,
		updatedNotice.DbContact, updatedNotice.DbDetails,
		updatedNotice.DbNoticeId)
	CheckError(e)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\nRows Updated = %d", rowsAffected)
}

func InsertNotice(newNotice DbNotice) {
	db := connectToDB()

	// close database
	defer db.Close()

	result, e := db.Exec(insertNewNoticeStmnt, newNotice.DbHeading, newNotice.DbPrice,

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

	result, e := db.Exec(closeNoticeStmt, noticeId)
	CheckError(e)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("\nRows Affected = %d", rowsAffected)

}

func FetchNotices(status string) []DbNotice {
	// open database
	db := connectToDB()

	// close database
	defer db.Close()

	rows, err := db.Query(FetchNoticesStmt, "ACTIVE")
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

		err = rows.Scan(&noticeId, &heading, &price, &category, &area_level_1,
			&area_level_2, &contact, &details)
		CheckError(err)
		currentNotice := DbNotice{DbNoticeId: noticeId, DbHeading: heading,
			DbPrice:    price,
			DbCategory: category, DbAreaLevel1: area_level_1,
			DbAreaLavel2: area_level_2, DbContact: contact, DbDetails: details}
		notices = append(notices, currentNotice)

		fmt.Println(noticeId, heading, price, category, area_level_1,
			area_level_2, contact, details)
	}

	CheckError(err)
	return notices

}

func connectToDB() *sql.DB {
	psqlconn :=
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
