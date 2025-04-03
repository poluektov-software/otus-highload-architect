package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (s *service) login(c *gin.Context) {
	fmt.Print("/login ")
}
