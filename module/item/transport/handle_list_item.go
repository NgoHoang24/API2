package todotrpt

import (
	Profilebiz "Project/module/item/biz"
	Profile "Project/module/item/model"
	Profilestorage "Project/module/item/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleListProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging Profile.DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		storage := Profilestorage.NewSQLStorage(db)
		var biz = Profilebiz.NewListProfileBiz(storage)
		result, err := biz.ListProfile(c.Request.Context(), nil, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}
