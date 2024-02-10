package cmd

import (
	"fmt"
	"os"

	"go-200lab-g09/common"
	"go-200lab-g09/middleware"
	ginitem "go-200lab-g09/module/item/transport/gin"

	ginupload "go-200lab-g09/module/upload/transport/gin"
	userstorage "go-200lab-g09/module/user/storage"
	ginuser "go-200lab-g09/module/user/transport/gin"
	"go-200lab-g09/plugin/sdkgorm"
	"go-200lab-g09/plugin/tokenprovider/jwt"
	"go-200lab-g09/plugin/uploadprovider"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("social-todo-list"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main.mysql", common.PluginDBMain)),
		goservice.WithInitRunnable(jwt.NewJWTProvider(common.PluginJWT)),
		goservice.WithInitRunnable(uploadprovider.NewR2Provider(common.PluginR2)),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start social TODO service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()
		serviceLogger := service.Logger("service")

		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.Use(middleware.Recover())

			db := service.MustGet(common.PluginDBMain).(*gorm.DB)

			authStore := userstorage.NewSQLStore(db)
			authMiddleware := middleware.RequiredAuth(authStore, service)

			engine.Static("/static", "./static")
			v1 := engine.Group("/v1")
			{
				v1.POST("/register", ginuser.Register(service))
				v1.POST("/login", ginuser.Login(service))
				v1.GET("/profile", authMiddleware, ginuser.Profile())

				uploads := v1.Group("/upload", authMiddleware)
				{
					// uploads.POST("", ginupload.Upload())
					uploads.POST("/local", ginupload.UploadLocal(service))
				}

				items := v1.Group("/items", authMiddleware)
				{
					items.POST("", ginitem.CreateItem(service))
					items.GET("", ginitem.ListItem(service))
					items.GET("/:id", ginitem.GetItem(service))
					items.PATCH("/:id", ginitem.UpdateItem(service))
					items.DELETE("/:id", ginitem.DeleteItem(service))
				}
			}
		})

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
