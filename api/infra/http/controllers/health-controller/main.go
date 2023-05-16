package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Alive(c *gin.Context) {
	log.Print("Alive: ", gin.H{"status": "alive"})
	c.JSON(http.StatusOK, gin.H{"status": "alive"})
}
