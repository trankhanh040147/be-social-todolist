package main

import (
	"social-todo-list/cmd"
)

func main() {
	cmd.Execute()
}

// RDS Service
// host: trankhanh-rds.c3kmq2oy4fy8.us-east-1.rds.amazonaws.com
// user: root
// password: Trankhanh47
// port: 3306
// database: social-todo-list

// MYSQL_GORM_DB_SOURCE_RDS = "root:Trankhanh47@tcp(trankhanh-rds.c3kmq2oy4fy8.us-east-1.rds.amazonaws.com:3306)/social-todo-list?charset=utf8mb4&parseTime=True&loc=Local"

// func main() {

// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
// 	// DB_CONNECTION from launch.json
// 	// dsn := os.Getenv("DB_CONNECTION")
// 	dsn := "root:my-secret-pw@tcp(127.0.0.1:3309)/social-todo-list?charset=utf8mb4&parseTime=True&loc=Local"
// 	// systemSecret := os.Getenv("SECRET")
// 	systemSecret := "iTaskSecret2024"

// 	db, err := gorm.Open(mysql.Open(string(dsn)), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	db = db.Debug()
// 	log.Println("DB Connection: ", db)

// 	// -----------------------------------------------
// 	authStore := storage.NewSQLStore(db)
// 	tokenProvider := jwt.NewTokenJwtProvider("jwt", systemSecret)
// 	middlewareAuth := middleware.RequiredAuth(authStore, tokenProvider)

// 	jsString := `{"id":1,"title":"Belajar Golang","description":"Belajar Golang untuk membuat API","status":"Active","created_at":"2021-10-13T15:04:05Z","updated_at":"2021-10-13T15:04:05Z"}`
// 	var item2 model.TodoItem
// 	if err := json.Unmarshal([]byte(jsString), &item2); err != nil {
// 		log.Fatalln(err)
// 	}
// 	log.Println(item2)

// 	// -----------------------------------------------

// 	// > set up HTTP routes for a RESTful API.
// 	// >> creates a new Gin router with default middleware. The default middleware includes logging and recovery middleware, which logs all requests and recovers from any panics, respectively.
// 	r := gin.Default()
// 	r.Use(middleware.Recover())

// 	r.Static("/static", "./static")
// 	// >> creates a new route group. All routes defined under this group will have the prefix /api/v1.
// 	v1 := r.Group("/api/v1")
// 	{
// 		// >> creates a new route handler /upload
// 		v1.PUT("/upload", upload.Upload(db))

// 		v1.POST("/register", ginuser.Register(db))
// 		v1.POST("/login", ginuser.Login(db, tokenProvider))
// 		v1.GET("/profile", middlewareAuth, ginuser.Profile())

// 		// >> all routes defined under this group will have the prefix /api/v1/items.
// 		items := v1.Group("items", middlewareAuth)
// 		{
// 			items.POST("", ginitem.CreateItem(db))
// 			items.GET("", ginitem.ListItem(db))
// 			items.GET("/:id", ginitem.GetItem(db))
// 			items.PATCH("/:id", ginitem.UpdateItem(db))
// 			items.DELETE("/:id", ginitem.DeleteItem((db)))
// 		}
// 	}

// 	// > define a route handler / that returns a JSON response with a message property set to pong.
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, common.SimpleSuccessResponse("pong"))
// 	})

// 	// > run the application on port 3000.
// 	if err := r.Run(":3000"); err != nil {
// 		log.Fatalln(err)
// 	}

// 	// fmt.Println("Hello, World!")
// 	// fmt.Println(os.Getenv("APP_NAME"))
// }
