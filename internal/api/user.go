package api

import (
	"cmd/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *service) create(c *gin.Context) {
	var (
		user model.User
		err  error
	)
	if err = c.BindJSON(&user); err != nil {
		fmt.Print("User data error")
		return
	}
	err = s.storage.AddUser(c.Request.Context(), model.Account{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	if err != nil {
		fmt.Print("User not added")
	}
}

func (s *service) get(c *gin.Context) {
	var (
		id  int
		err error
	)
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Print("User id invalid")
		c.Status(http.StatusBadRequest)
		return
	}
	user, err := s.storage.GetUser(c.Request.Context(), id)
	if err != nil {
		fmt.Print("User not found")
		c.Status(http.StatusNotFound)
		return
	}
	fmt.Printf("first_name = %s, last_name = %s", user.FirstName, user.LastName)
	c.JSON(http.StatusOK, user)
}
