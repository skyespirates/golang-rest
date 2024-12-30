package controllers

import "github.com/gin-gonic/gin"

type Data struct {
	Message string `json:"msg"`
}

func GetAllBook(ctx *gin.Context) {
	res := Data{Message: "Retrieve all book"}
	ctx.JSON(200, res)
}

func GetSingleBook(ctx *gin.Context) {
	res := Data{Message: "Retrieve single book"}
	ctx.JSON(200, res)
}