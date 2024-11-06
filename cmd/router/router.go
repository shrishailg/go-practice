package router

/*
40 times faster than Martini

*/

import (
	"github.com/gin-gonic/gin"
	"github.com/go-practice/cmd/middleware"
	"github.com/go-practice/cmd/server"
	"github.com/go-practice/internal/handlers"
)

var (
	requestHandlers *handlers.RequestHandler
)

func Init() {

	//below 2 methods are same, router.run calls http.listen and serve method
	// router.Run(":8080")

	// http.ListenAndServe(":9090", router)

	router := NewRouter()

	server := server.NewServer(router)

	server.ListenAndServe()
}

func NewRouter() *gin.Engine {
	// this will create gin with default middleware for logger and recovery
	// router := gin.Default()

	// we can use below method also but it cant attach default logger and reovery mechanism

	//to add default logger
	// router.Use(gin.Logger())

	router := gin.New()

	// gin context contains all the info about request and response
	// router.GET("/get/data", handleGetData)

	router.GET("/get/data/:name/:age", requestHandlers.HandlePathParamGetData)

	// router.GET("/getQueryString", handleQueryGetData)

	router.POST("/createData", middleware.SetResponseHeader, requestHandlers.HandleDataCreation)

	initAdminRouters(router)
	initClientRouters(router)

	return router
}

func initAdminRouters(router *gin.Engine) {
	//route grouping
	admin := router.Group("/admin")
	{
		admin.GET("/get/data", requestHandlers.HandleGetData)
	}
}

func initClientRouters(router *gin.Engine) {
	client := router.Group("/client", middleware.Authentication, middleware.SetResponseHeader)
	{
		client.GET("/getQueryString", requestHandlers.HandleQueryGetData)
	}
}
