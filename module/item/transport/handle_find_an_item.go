package todotrpt

import (
	Profilebiz "Project/module/item/biz"
	Profilestorage "Project/module/item/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleFindAProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := Profilestorage.NewSQLStorage(db)
		biz := Profilebiz.NewFindProfileBiz(storage)

		data, err := biz.FindAProfile(c.Request.Context(), map[string]interface{}{"id": id})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
