package nbNet

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swathiGiligar/nbServer/nbDB"
)

type notice struct {
	NoticeID   string    `json:"notice_id"`
	Heading    string    `json:"heading"`
	Price      string    `json:"price"`
	Category   string    `json:"category"`
	AreaLevel1 string    `json:"area_level_1"`
	AreaLavel2 string    `json:"area_level_2"`
	Contact    string    `json:"contact"`
	Details    string    `json:"details"`
	Created_On time.Time `json:"created_on"`
}

func SetRouter() {
	router := gin.Default()
	router.GET("/noticeBoard", GetNotices)
	router.POST("/noticeBoard", AddNotices)
	router.PATCH("/noticeBoard/:notice_id", CloseNotice)
	router.PUT("/noticeBoard", UpdateNotice)

	router.Run(":8080")
}

func GetNotices(c *gin.Context) {
	var notices = nbDB.FetchNotices("ACTIVE")
	var jsonNotices = []notice{}

	for _, current := range notices {

		jsonNotice := notice{NoticeID: strconv.FormatInt(current.DbNoticeId, 10),
			Heading: current.DbHeading, Price: current.DbPrice,
			Category: current.DbCategory, AreaLevel1: current.DbAreaLevel1,
			AreaLavel2: current.DbAreaLavel2, Contact: current.DbContact,
			Details: current.DbDetails, Created_On: current.DbCreatedOn}
		jsonNotices = append(jsonNotices, jsonNotice)
	}

	c.IndentedJSON(http.StatusOK, jsonNotices)
}

func AddNotices(c *gin.Context) {
	var newNotice notice

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newNotice); err != nil {
		panic(err)
	}

	toDB := nbDB.DbNotice{DbHeading: newNotice.Heading, DbPrice: newNotice.Price,
		DbCategory: newNotice.Category, DbAreaLevel1: newNotice.AreaLevel1,
		DbAreaLavel2: newNotice.AreaLavel2, DbContact: newNotice.Contact,
		DbDetails: newNotice.Details}

	nbDB.InsertNotice(toDB)
	c.IndentedJSON(http.StatusCreated, newNotice)
}

func CloseNotice(c *gin.Context) {
	noticeId := c.Param("notice_id")
	noticeIdNumber, err := strconv.ParseInt(noticeId, 10, 64)
	if err == nil {
		nbDB.CloseNotice(noticeIdNumber)
	}
	c.IndentedJSON(http.StatusOK, noticeId)
}

func UpdateNotice(c *gin.Context) {
	var updatedNotice notice

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&updatedNotice); err != nil {
		panic(err)
	}

	noticeIdNumber, errParse := strconv.ParseInt(updatedNotice.NoticeID, 10, 64)
	if errParse != nil {
		panic(errParse)
	}

	toDB := nbDB.DbNotice{DbNoticeId: noticeIdNumber, DbHeading: updatedNotice.Heading,
		DbPrice: updatedNotice.Price, DbCategory: updatedNotice.Category,
		DbAreaLevel1: updatedNotice.AreaLevel1,
		DbAreaLavel2: updatedNotice.AreaLavel2, DbContact: updatedNotice.Contact,
		DbDetails: updatedNotice.Details}

	nbDB.UpdateNotice(toDB)
	c.IndentedJSON(http.StatusOK, updatedNotice)
}
