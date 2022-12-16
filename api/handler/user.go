package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mokh1rbek/golang_CRUD/models"
	"github.com/mokh1rbek/golang_CRUD/storage"
)

// CreateUser godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "CreateUserRequestBody"
// @Success 201 {object} models.User "GetUserBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) Create(c *gin.Context) {

	var (
		user models.User
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := storage.Create(h.db, user)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling create").Error())
		return
	}

	userId, err := storage.GetById(h.db, id)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusCreated, userId)
}

// GetByIdUser godoc
// @ID get_by_id_user
// @Router /user/{id} [GET]
// @Summary Get By Id User
// @Description Get By Id User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.User "GetUserBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetById(c *gin.Context) {

	id := c.Param("id")

	user, err := storage.GetById(h.db, id)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetListUser godoc
// @ID get_list_user
// @Router /user [GET]
// @Summary Get List User
// @Description Get List User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "CreateUserRequestBody"
// @Success 201 {object} models.User "GetUserBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetList(c *gin.Context) {

	users, err := storage.GetList(h.db)
	if err != nil {
		log.Printf("error whiling get list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get list").Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *HandlerV1) Update(c *gin.Context) {

	var (
		user models.User
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := storage.Update(h.db, user)
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

	resp, err := storage.GetById(h.db, user.Id)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *HandlerV1) Patch(c *gin.Context) {

	var (
		user models.UserData
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Printf("error whiling patch: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := storage.Patch(h.db, user)
	if err != nil {
		log.Printf("error whiling patch: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling patch").Error())
		return
	}

	if rowsAffected == 0 {
		log.Printf("error whiling update rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling patch rows affected").Error())
		return
	}

	resp, err := storage.GetById(h.db, user.Id)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *HandlerV1) Delete(c *gin.Context) {

	id := c.Param("id")

	err := storage.Delete(h.db, id)
	if err != nil {
		log.Printf("error whiling delete: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling delete").Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
