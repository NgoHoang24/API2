package todotrpt

import (
	Profilebiz "Project/module/item/biz"
	Profile "Project/module/item/model"
	Profilestorage "Project/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func HandleCreateProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem Profile.Profile

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// setup dependencies
		storage := Profilestorage.NewSQLStorage(db)
		biz := Profilebiz.NewCreateProfileBiz(storage)

		if err := biz.CreateNewProfile(c.Request.Context(), &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem.Id})
	}
}
