package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mokh1rbek/golang_CRUD/api/docs"
	"github.com/mokh1rbek/golang_CRUD/api/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpApi(r *gin.Engine, db *sql.DB) {

	handlerV1 := handler.NewHandlerV1(db)

	r.POST("/film", handlerV1.CreateFilm)
	r.GET("/film/:id", handlerV1.GetFilmById)
	r.GET("/film", handlerV1.GetFilmList)
	r.PUT("/film", handlerV1.UpdateFilm)
	r.DELETE("/film/:id", handlerV1.DeleteFilm)

	r.POST("/actor", handlerV1.CreateActor)
	r.GET("/actor/:id", handlerV1.GetActorById)
	r.GET("/actor", handlerV1.GetActorList)
	r.PUT("/actor", handlerV1.UpdateActor)
	r.DELETE("/actor/:id", handlerV1.DeleteActor)

	r.POST("/category", handlerV1.CreateCategory)
	r.GET("/category/:id", handlerV1.GetCategoryById)
	r.GET("/category", handlerV1.GetCategoryList)
	r.PUT("/category", handlerV1.UpdateCategory)
	r.DELETE("/category/:id", handlerV1.DeleteCategory)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
