package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mokh1rbek/golang_CRUD/models"
	"github.com/mokh1rbek/golang_CRUD/storage"
)

// CreateCategory godoc
// @ID create_category
// @Router /category [POST]
// @Summary Create Category
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param category body models.FilmCategory true "CreateCategoryRequestBody"
// @Success 201 {object} models.FilmCategory "GetCategoryBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) CreateCategory(c *gin.Context) {

	var (
		category models.FilmCategory
	)

	err := c.ShouldBindJSON(&category)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := storage.CreateCategory(h.db, category)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling create").Error())
		return
	}

	CategoryId, err := storage.GetCategoryById(h.db, id)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusCreated, CategoryId)
}

// GetByIdCategory godoc
// @ID get_by_id_category
// @Router /category/{id} [GET]
// @Summary Get By Id Category
// @Description Get By Id Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.FilmCategory "GetCategoryBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetCategoryById(c *gin.Context) {

	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	category, err := storage.GetCategoryById(h.db, newId)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusOK, category)
}

// GetListCategory godoc
// @ID get_list_category
// @Router /category [GET]
// @Summary Get List Category
// @Description Get List Category
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} []models.FilmCategory "GetCategoryBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetCategoryList(c *gin.Context) {

	categories, err := storage.GetCategoryList(h.db)
	if err != nil {
		log.Printf("error whiling get list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get list").Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}

// UpdateCategory godoc
// @ID update_category
// @Router /category/ [PUT]
// @Summary Update Category
// @Description Update Category
// @Tags Category
// @Accept json
// @Produce json
// @Param user body models.FilmCategory true "CreateCategoryRequestBody"
// @Success 200 {object} models.FilmCategory "GetCategoriesBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) UpdateCategory(c *gin.Context) {

	var (
		category models.FilmCategory
	)

	err := c.ShouldBindJSON(&category)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := storage.UpdateCategory(h.db, category)
	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	fmt.Println(rowsAffected)

	if rowsAffected == 0 {
		log.Printf("error whiling update rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update rows affected").Error())
		return
	}

	resp, err := storage.GetCategoryById(h.db, category.CategoryId)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteByIdCategory godoc
// @ID delete_by_id_category
// @Router /category/{id} [DELETE]
// @Summary Delete By Id Category
// @Description Delete By Id Category
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.FilmCategory "GetCategoryBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) DeleteCategory(c *gin.Context) {

	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	err = storage.DeleteCategory(h.db, newId)
	if err != nil {
		log.Printf("error whiling delete: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling delete").Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
