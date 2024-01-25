package main

import (
	"demo-mod/common"
	"encoding/json"
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

	// item := TodoItem{
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
	var item2 TodoItem
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
			items.POST("", CreateItem(db))
			items.GET("", ListItem(db))
			items.GET("/:id", GetItem(db))
			items.PATCH("/:id", UpdateItem(db))
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

// > create a handler for: POST /api/v1/items
// func CreateItem() gin.HandlerFunc {
// CreateItem is a handler function that creates a new todo item in the database.
// It takes a *gorm.DB as a parameter and returns a gin.HandlerFunc.
// The function parses JSON from the HTTP request body to a TodoItem struct,
// validates the data, and returns an error response if the data is invalid.
// If the data is valid, it creates a new todo item in the database.
func CreateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		// step 1: Parse JSON from HTTP request body to TodoItem struct
		var itemData TodoItemCreation

		// step 2: Validate the data
		// >> The ShouldBind method automatically infers the content type and binds the data into the type of the provided struct.
		// >> The & operator before itemData is used to pass the memory address of itemData to the ShouldBind method. This means any changes made to itemData inside the ShouldBind method will affect the original itemData variable.
		// >> If ShouldBind encounters an error during the binding process (for example, if the provided JSON does not match the structure of itemData), it will return a non-nil error. The err != nil check will then evaluate to true and the code inside the if block will be executed.
		if err := c.ShouldBind(&itemData); err != nil {
			// >> c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),}) is called. This sends a JSON response to the client with a status code of 400 (Bad Request). The JSON body of the response contains a single property error, which is set to the string representation of the error returned by ShouldBind. The gin.H function is a shortcut for creating a map in Go, which in this case is used to create the JSON object.
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// step 3: use db.Create to
		if err := db.Create(&itemData).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// step 4: print data of the inserted record
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(itemData))
	}
}

func GetItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var itemData TodoItem

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", id).First(&itemData).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(itemData))
	}
}

func UpdateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var updateData TodoItem

		if err := c.ShouldBind(&updateData); err != nil {
			// >> c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),}) is called. This sends a JSON response to the client with a status code of 400 (Bad Request). The JSON body of the response contains a single property error, which is set to the string representation of the error returned by ShouldBind. The gin.H function is a shortcut for creating a map in Go, which in this case is used to create the JSON object.
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Where("id = ?", id).Updates(&updateData).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(updateData))
	}
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

		if err := db.Where("id = ?", id).Updates(&TodoItemUpdate{Status: &deletedStatus}).Error; err != nil {
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

		var result []TodoItem

		db = db.Table(TodoItem{}.TableName()).Where("status <> ?", "Deleted")

		// SELECT COUNT(id) FROM todo_items --> paging.Total
		if err := db.Table(TodoItem{}.TableName()).Select("id").Count(&paging.Total).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		// SELECT * FROM todo_items ORDER BY id ASC LIMIT paging.Limit OFFSET (paging.Page - 1) * paging.Limit
		if err := db.Table(TodoItem{}.TableName()).
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
