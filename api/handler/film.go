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

// CreateFilm godoc
// @ID create_film
// @Router /film [POST]
// @Summary Create Film
// @Description Create Film
// @Tags Film
// @Accept json
// @Produce json
// @Param film body models.Film true "CreateFilmRequestBody"
// @Success 201 {object} models.Film "GetFilmBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) CreateFilm(c *gin.Context) {

	var (
		film models.Film
	)

	err := c.ShouldBindJSON(&film)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := storage.CreateFilm(h.db, film)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling create").Error())
		return
	}

	filmId, err := storage.GetFilmById(h.db, id)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusCreated, filmId)
}

// GetByIdFilm godoc
// @ID get_by_id_film
// @Router /film/{id} [GET]
// @Summary Get By Id Film
// @Description Get By Id Film
// @Tags Film
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Film "GetFilmBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetFilmById(c *gin.Context) {

	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	film, err := storage.GetFilmById(h.db, int(newId))
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusOK, film)
}

// GetListFilm godoc
// @ID get_list_film
// @Router /film [GET]
// @Summary Get List Film
// @Description Get List Film
// @Tags Film
// @Accept json
// @Produce json
// @Success 200 {object} []models.Film "GetFilmBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetFilmList(c *gin.Context) {

	films, err := storage.GetFilmList(h.db)
	if err != nil {
		log.Printf("error whiling get list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get list").Error())
		return
	}

	c.JSON(http.StatusOK, films)
}

// UpdateFilm godoc
// @ID update_film
// @Router /film/ [PUT]
// @Summary Update Film
// @Description Update Film
// @Tags Film
// @Accept json
// @Produce json
// @Param film body models.Film true "CreateFilmRequestBody"
// @Success 200 {object} models.Film "GetFilmsBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) UpdateFilm(c *gin.Context) {

	var (
		film models.Film
	)

	err := c.ShouldBindJSON(&film)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := storage.UpdateFilm(h.db, film)
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

	resp, err := storage.GetFilmById(h.db, film.FilmId)
	if err != nil {
		log.Printf("error whiling get by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get by id").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteByIdFilm godoc
// @ID delete_by_id_film
// @Router /film/{id} [DELETE]
// @Summary Delete By Id Film
// @Description Delete By Id Film
// @Tags Film
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Film "GetFilmBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) DeleteFilm(c *gin.Context) {

	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	err = storage.DeleteFilm(h.db, newId)
	if err != nil {
		log.Printf("error whiling delete: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling delete").Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
