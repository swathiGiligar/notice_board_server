package resources

const (

	//Notice Management

	FetchNoticesStmt = `SELECT "notice_id", "heading", "price", "category",
	"area_level_1", "area_level_2", "contact", "details", "created_on" 
	 FROM "notices" WHERE status=$1 order by updated_on desc`

	InsertNewNoticeStmnt = `insert into "notices"("heading", "price",
	 "category", "area_level_1", "area_level_2", "contact", "details", 
	 "status") values($1, $2, $3, $4, $5, $6, $7, $8)`

	CloseNoticeStmt = `update "notices" set "status"='CLOSED', "updated_on"=$1  
	where "notice_id"=$2`

	UpdateNoticeStmt = `update "notices" set "heading"=$1, "price"=$2,
	"category"=$3, "area_level_1"=$4, "area_level_2"=$5, "contact"=$6, 
	"details"=$7, "updated_on"=$8  where "notice_id"=$9`

	//User Management

	FetchUserWithLoginIdStmt = `select "user_id", "user_name", "login_id",
	 "access_level", "last_login" from "USERS" where "login_id" = $1;`

	InsertNewUserStmt = `insert into "USERS"("user_name", "login_id", 
	"access_level", "last_login") values ($1, $2, $3, current_timestamp);`

	UpdateLastLoginStmt = `update "users" set "last_login" = current_timestamp 
	where "user_id" = $1;`

	UpdateUserAccessStmt = `update "users" set "access_level"  = $1
	 where "user_id" = $2;`
)
