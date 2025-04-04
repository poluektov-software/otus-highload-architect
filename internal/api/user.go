package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"otus-highload-architect/internal/model"
	"strconv"
)

func (s *service) create(c *gin.Context) {
	var (
		user model.User
		err  error
	)
	c.Header("Content-Type", "application/json")
	if err = c.ShouldBindJSON(&user); err != nil {
		log.Printf("User data error: %s", err.Error())
		c.Status(http.StatusUnprocessableEntity)
		return
	}
	err = s.storage.AddUser(c.Request.Context(), model.Account{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	if err != nil {
		log.Printf("User not added: %s", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}
}

func (s *service) get(c *gin.Context) {
	var (
		id  int
		err error
	)
	c.Header("Content-Type", "application/json")
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Print("User id invalid")
		c.Status(http.StatusBadRequest)
		return
	}
	user, err := s.storage.GetUser(c.Request.Context(), id)
	if err != nil {
		log.Print("User not found")
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}
