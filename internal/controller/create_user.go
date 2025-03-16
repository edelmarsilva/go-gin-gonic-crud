package controller

import (
	"crud-go/internal/controller/model/request"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(userRequest)

}
