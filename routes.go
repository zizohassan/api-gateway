package main

import (
	"github.com/gin-gonic/gin"
)

// here we set up all routes
// add middleware and return with gin instance to main file
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	// routes in this micro services without any limitation
	r.GET("/", Home)
	// auth area in this micro service
	auth := r.Group("auth")
	auth.Use(AuthRequired())
	{
		auth.GET("/profile", Profile)
	}
	// route in other micro service
	// url must have the name of micro service
	r.Use(handelOtherMicroServiceRequests())
	return r
}
