package todotrpt

import (
	Profilebiz "Project/module/item/biz"
	Profilestorage "Project/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func HandleDeleteAProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// setup dependencies
		storage := Profilestorage.NewSQLStorage(db)
		biz := Profilebiz.NewDeleteProfileBiz(storage)

		if err := biz.DeleteProfile(c.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
