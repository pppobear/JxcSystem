package staff

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/filters/auth"
)

func SetGroup(rootGroup *gin.RouterGroup) {
	staffGroup := rootGroup.Group("/staff")
	{
		staffGroup.POST("", CreateStaff)
		staffGroup.GET("", ListStaff).Use(auth.Middleware("jwt"))
		staffGroup.GET("/:id", RetrieveStaff).Use(auth.Middleware("jwt"))
		staffGroup.PUT("/:id", UpdateStaff).Use(auth.Middleware("jwt"))
		staffGroup.DELETE("/:id", DeleteStaff).Use(auth.Middleware("jwt"))
	}

	specialtyGroup := rootGroup.Group("/specialty")
	{
		specialtyGroup.GET("", ListSpe)
		specialtyGroup.POST("", CreateSpe).Use(auth.Middleware("jwt"))
		specialtyGroup.GET("/:id", RetrieveSpe).Use(auth.Middleware("jwt"))
		specialtyGroup.PUT("/:id", UpdateSpe).Use(auth.Middleware("jwt"))
		specialtyGroup.DELETE("/:id", DeleteSpe).Use(auth.Middleware("jwt"))
	}

	departmentGroup := rootGroup.Group("/department")
	{
		departmentGroup.GET("", ListDep)
		departmentGroup.POST("", CreateDep).Use(auth.Middleware("jwt"))
		departmentGroup.GET("/:id", RetrieveDep).Use(auth.Middleware("jwt"))
		departmentGroup.PUT("/:id", UpdateDep).Use(auth.Middleware("jwt"))
		departmentGroup.DELETE("/:id", DeleteDep).Use(auth.Middleware("jwt"))
	}
}
