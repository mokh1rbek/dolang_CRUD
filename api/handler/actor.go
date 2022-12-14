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

// CreateActor godoc
// @ID create_actor
// @Router /actor [POST]
// @Summary Create Actor
// @Description Create Actor
// @Tags Actor
// @Accept json
// @Produce json
// @Param actor body models.FilmActor true "CreateActorRequestBody"
// @Success 201 {object} models. "GetUserBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) CreateActor(c *gin.Context) {

	var (
		actor models.FilmActor
	)

	err := c.ShouldBindJSON(&actor)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := storage.CreateActor(h.db, actor)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling create").Error())
		return
	}

	actorId, err := storage.GetActorById(h.db, id)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusCreated, actorId)
}

// GetByIdActor godoc
// @ID get_by_id_actor
// @Router /actor/{id} [GET]
// @Summary Get By Id Actor
// @Description Get By Id Actor
// @Tags Actor
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.FilmActor "GetActorBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetActorById(c *gin.Context) {

	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	actor, err := storage.GetActorById(h.db, newId)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusOK, actor)
}

// GetListActor godoc
// @ID get_list_actor
// @Router /actor [GET]
// @Summary Get List Actor
// @Description Get List Actor
// @Tags Actor
// @Accept json
// @Produce json
// @Success 200 {object} []models.FilmActor "GetActorBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetActorList(c *gin.Context) {

	actors, err := storage.GetActorList(h.db)
	if err != nil {
		log.Printf("error whiling get list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get list").Error())
		return
	}

	c.JSON(http.StatusOK, actors)
}

// UpdateActor godoc
// @ID update_actor
// @Router /actor/ [PUT]
// @Summary Update Actor
// @Description Update Actor
// @Tags Actor
// @Accept json
// @Produce json
// @Param user body models.FilmActor true "CreateFilmRequestBody"
// @Success 200 {object} models.FilmActor "GetFilmsBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) UpdateActor(c *gin.Context) {

	var (
		actor models.FilmActor
	)

	err := c.ShouldBindJSON(&actor)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := storage.UpdateActor(h.db, actor)
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

	resp, err := storage.GetActorById(h.db, actor.ActorId)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteByIdActor godoc
// @ID delete_by_id_actor
// @Router /actor/{id} [DELETE]
// @Summary Delete By Id Actor
// @Description Delete By Id Actor
// @Tags Actor
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.FilmActor "GetActorBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) DeleteActor(c *gin.Context) {

	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	err = storage.DeleteActor(h.db, newId)
	if err != nil {
		log.Printf("error whiling delete: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling delete").Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
