package upload

import (
	"fmt"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
	"social-todo-list/common"
	"social-todo-list/module/upload/biz"
	"social-todo-list/module/upload/storage"
	"social-todo-list/plugin/uploadprovider"
	"time"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
)

func UploadLocal(serviceCtx goservice.ServiceContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		// db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		dst := fmt.Sprintf("static/%d.%s", time.Now().UnixNano(), fileHeader.Filename)

		if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
			// panic(common.ErrInvalidRequest(err))
		}

		img := common.Image{
			Id:        0,
			Url:       dst,
			Width:     100,
			Height:    100,
			CloudName: "local",
			Extension: "",
		}

		img.FullFill("http://localhost:3000")

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}

func Upload(serviceCtx goservice.ServiceContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		provider := serviceCtx.MustGet(common.PluginS3).(uploadprovider.UploadProvider)

		_, dataBytes, folder, fileName, contentType := validateFiles(c)
		store := storage.NewSQLStore(db)
		business := biz.NewUploadBiz(store, provider)

		img, err := business.Upload(c.Request.Context(), dataBytes, folder, fileName, contentType)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}

func validateFiles(ctx *gin.Context) (fileHeader *multipart.FileHeader, dataBytes []byte, folder, fileName, contentType string) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		panic(common.ErrInvalidRequest(err))
	}

	folder = ctx.DefaultPostForm("folder", "images")
	file, err := fileHeader.Open()
	if err != nil {
		panic(common.ErrInvalidRequest(err))
	}

	defer file.Close()

	fileName = fileHeader.Filename
	contentType = fileHeader.Header.Get("Content-Type")
	dataBytes = make([]byte, fileHeader.Size)
	if _, err := file.Read(dataBytes); err != nil {
		panic(common.ErrInvalidRequest(err))
	}

	return
}
