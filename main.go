package main

import (
	"encoding/json"
	"go-200lab-g09/common"
	"go-200lab-g09/module/item/model"
	ginitem "go-200lab-g09/module/item/transport/gin"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// DB_CONNECTION from launch.json
	// dsn := os.Getenv("DB_CONNECTION")
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3309)/social-todo-list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(string(dsn)), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db = db.Debug()
	log.Println("DB Connection: ", db)

	// now := time.Now().UTC()

	// item := model.TodoItem{
	// 	Id:          1,
	// 	Title:       "Belajar Golang",
	// 	Description: "Belajar Golang untuk membuat API",
	// 	Status:      "Active",
	// 	CreatedAt:   &now,
	// 	UpdatedAt:   &now,
	// }

	// jsData, err := json.Marshal(item)

	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(string(jsData))

	jsString := `{"id":1,"title":"Belajar Golang","description":"Belajar Golang untuk membuat API","status":"Active","created_at":"2021-10-13T15:04:05Z","updated_at":"2021-10-13T15:04:05Z"}`
	var item2 model.TodoItem
	if err := json.Unmarshal([]byte(jsString), &item2); err != nil {
		log.Fatalln(err)
	}
	log.Println(item2)

	// -----------------------------------------------

	// > set up HTTP routes for a RESTful API.
	// >> creates a new Gin router with default middleware. The default middleware includes logging and recovery middleware, which logs all requests and recovers from any panics, respectively.
	r := gin.Default()
	// >> creates a new route group. All routes defined under this group will have the prefix /api/v1.
	v1 := r.Group("/api/v1")
	{
		// >> all routes defined under this group will have the prefix /api/v1/items.
		items := v1.Group("items")
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", DeleteItem((db)))
		}
	}

	// > define a route handler / that returns a JSON response with a message property set to pong.
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.SimpleSuccessResponse("pong"))
	})

	// > run the application on port 3000.
	if err := r.Run(":3000"); err != nil {
		log.Fatalln(err)
	}

	// fmt.Println("Hello, World!")
	// fmt.Println(os.Getenv("APP_NAME"))
}

func DeleteItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		deletedStatus := "Deleted"

		if err := db.Where("id = ?", id).Updates(&model.TodoItemUpdate{Status: &deletedStatus}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

func ListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()

		var result []model.TodoItem

		db = db.Table(model.TodoItem{}.TableName()).Where("status <> ?", "Deleted")

		// SELECT COUNT(id) FROM todo_items --> paging.Total
		if err := db.Table(model.TodoItem{}.TableName()).Select("id").Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		// SELECT * FROM todo_items ORDER BY id ASC LIMIT paging.Limit OFFSET (paging.Page - 1) * paging.Limit
		if err := db.Table(model.TodoItem{}.TableName()).
			Select("*").
			Offset((paging.Page - 1) * paging.Limit).
			Limit(paging.Limit).
			Order("id asc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
