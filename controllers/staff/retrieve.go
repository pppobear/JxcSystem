package staff

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

func RetrieveStaff(c *gin.Context) {
	id := c.Param("id")
	staff, err := models.RetrieveStaff(id)
	if err != nil {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, staff)
}

func RetrieveSpe(c *gin.Context) {
	id := c.Param("id")
	specialty, err := models.RetrieveSpecialty(id)
	if err != nil {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, specialty)
}

func RetrieveDep(c *gin.Context) {
	id := c.Param("id")
	department, err := models.RetrieveDepartment(id)
	if err != nil {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, department)
}
