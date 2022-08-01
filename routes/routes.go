package routes

import (
	"Kuisioner-MySql/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/user", controllers.GetUser)
	r.POST("/user", controllers.TambahUser)
	r.GET("/user/:id", controllers.CariUser)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.DELETE("user/:id", controllers.DeleteUser)

	r.GET("/kuisioner", controllers.GetKuis)
	r.POST("/kuisioner", controllers.TambahKuis)
	r.GET("/search/:isi", controllers.Search)
	r.GET("/kuisioner/:id", controllers.CariKuis)
	r.PATCH("/kuisioner/:id", controllers.UpdateKuis)
	r.DELETE("kuisioner/:id", controllers.DeleteKuis)
	return r
}
