package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type StashRequestBody struct {
	Database   string      `json:"database"`
	Collection string      `json:"collection"`
	Document   interface{} `json:"document"`
}

func stash(c *gin.Context) {
	var json StashRequestBody
	if err := c.ShouldBindWith(&json, binding.JSON); err == nil {
		go func() {
			_, err := insertMongoDocument(json.Database, json.Collection, json.Document)
			if err != nil {
				log.Error(err.Error())
			}
		}()
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
