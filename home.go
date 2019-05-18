package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// this function any one can access it
func Home(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "Welcome in Api Get way",
		"status":  true,
		"data":    "",
	})
}

// this function run after auth middleware
func Profile(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "You ara login",
		"status":  true,
		"data":    "",
	})
}


