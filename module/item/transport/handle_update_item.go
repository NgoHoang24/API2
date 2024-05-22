package todotrpt

import (
	Profilebiz "Project/module/item/biz"
	Profile "Project/module/item/model"
	Profilestorage "Project/module/item/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleUpdateAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dataItem Profile.Profile

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := Profilestorage.NewSQLStorage(db)
		biz := Profilebiz.NewUpdateProfileBiz(storage)

		if err := biz.UpdateProfile(c.Request.Context(), map[string]interface{}{"id": id}, &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
