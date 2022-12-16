package api

import (
	_ "github.com/mokh1rbek/CRUD/api/docs"
	"github.com/mokh1rbek/CRUD/api/handler"
	"github.com/mokh1rbek/CRUD/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpApi(r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandlerV1(storage)

	r.POST("/film", handlerV1.CreateFilm)
	r.GET("/film/:id", handlerV1.GetFilmById)
	r.GET("/film", handlerV1.GetFilmList)
	r.PUT("/film/:id", handlerV1.UpdateFilm)
	r.DELETE("/film/:id", handlerV1.DeleteFilm)

	r.POST("/category", handlerV1.CreateCategory)
	r.GET("/category/:id", handlerV1.GetCategoryById)
	r.GET("/category", handlerV1.GetCategoryList)
	r.PUT("/category/:id", handlerV1.UpdateCategory)
	r.DELETE("/category/:id", handlerV1.DeleteCategory)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
