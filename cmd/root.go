package cmd

import (
	"fmt"
	"log"
	"os"
	rpcuserlikeitem "social-todo-list/module/userlikeitem/transport/rpc"
	"social-todo-list/plugin/rpccaller"
	"social-todo-list/plugin/simple"
	"social-todo-list/plugin/uploadprovider"
	"social-todo-list/pubsub"
	subscriber "social-todo-list/subcriber"

	"social-todo-list/common"
	"social-todo-list/middleware"
	ginitem "social-todo-list/module/item/transport/gin"
	ginuserlikeitem "social-todo-list/module/userlikeitem/transport"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	ginupload "social-todo-list/module/upload/transport/gin"
	userstorage "social-todo-list/module/user/storage"
	ginuser "social-todo-list/module/user/transport/gin"
	"social-todo-list/plugin/sdkgorm"
	"social-todo-list/plugin/tokenprovider/jwt"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("social-todo-list"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main.mysql", common.PluginDBMain)),
		goservice.WithInitRunnable(jwt.NewJWTProvider(common.PluginJWT)),
		//goservice.WithInitRunnable(uploadprovider.NewR2Provider(common.PluginR2)),
		goservice.WithInitRunnable(uploadprovider.NewS3Provider(common.PluginS3)),
		goservice.WithInitRunnable(simple.NewSimplePlugin("simple")),
		goservice.WithInitRunnable(rpccaller.NewApiItemCaller(common.PluginItemAPI)),
    goservice.WithInitRunnable(pubsub.NewPubSub(common.PluginPubSub)),

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

			log.Println(service.MustGet("simple").(interface {
				GetValue() string
			}).GetValue())

			db := service.MustGet(common.PluginDBMain).(*gorm.DB)

			authStore := userstorage.NewSQLStore(db)
			authMiddleware := middleware.RequiredAuth(authStore, service)

			engine.Static("/static", "./static")
			v1 := engine.Group("/api/v1")
			{
				v1.POST("/register", ginuser.Register(service))
				v1.POST("/login", ginuser.Login(service))
				v1.GET("/profile", authMiddleware, ginuser.Profile())

				uploads := v1.Group("/upload")
				{
					uploads.POST("", ginupload.Upload(service))
					uploads.POST("/local", ginupload.UploadLocal(service))
				}

				items := v1.Group("/items", authMiddleware)
				{
					items.POST("", ginitem.CreateItem(service))
					items.GET("", ginitem.ListItem(service))
					items.GET("/:id", ginitem.GetItem(service))
					items.PATCH("/:id", ginitem.UpdateItem(service))
					items.DELETE("/:id", ginitem.DeleteItem(service))

					items.POST("/:id/like", ginuserlikeitem.LikeItem(service))
					items.DELETE("/:id/unlike", ginuserlikeitem.UnlikeItem(service))
					items.GET("/:id/liked-users", ginuserlikeitem.ListLikedUsers(service))
				}

				rpc := v1.Group("/rpc")
				{
					rpc.POST("/get_item_likes", rpcuserlikeitem.GetItemLikes(service))
				}

			}
		})

		_ = subscriber.NewPBEngine(service).Start()

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
