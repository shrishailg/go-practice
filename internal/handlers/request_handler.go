package handlers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{}
}

func (r *RequestHandler) HandleGetData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "hellow there",
	})
}

// getQueryString/Mark/30
func (r *RequestHandler) HandlePathParamGetData(ctx *gin.Context) {
	name := ctx.Param("name")
	age := ctx.Param("age")
	ctx.JSON(http.StatusOK, gin.H{
		"data": "get path data",
		"name": name,
		"age":  age,
	})
}

// getQueryString?name=Mark&age=30
func (r *RequestHandler) HandleQueryGetData(ctx *gin.Context) {
	name := ctx.Query("name")
	age := ctx.Query("age")

	ctx.JSON(http.StatusOK, gin.H{
		"data": "get query data",
		"name": name,
		"age":  age,
	})
}

func (r *RequestHandler) HandleDataCreation(ctx *gin.Context) {
	requestBody := ctx.Request.Body
	if value, err := io.ReadAll(requestBody); err != nil {

		ctx.AbortWithStatusJSON(http.StatusCreated, gin.H{
			"error": "Error occured while reading the request body",
		})
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{
			"intro": "hellow there",
			"data":  string(value),
		})
	}

}
