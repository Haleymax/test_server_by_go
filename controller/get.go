package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test_server/model"
)

func List(c *gin.Context) {
	i := model.OrderList{}
	err, is := i.List()
	if err != nil {
		fmt.Println(err)
		c.JSON(200, nil)
	}
	c.JSON(200, is)
}
