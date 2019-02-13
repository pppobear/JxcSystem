package user

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"

	"pppobear.cn/jxc-backend/filters/auth"
	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

type LoginRequest struct {
	Id       string `gorm:"primary_key" json:"id"`
	Username string `gorm:"column:name" json:"username" binding:"required" validate:"min=1,max=8"`
	Password string `gorm:"column:password" json:"password" binding:"required" validate:"min=1,max=50"`
}

func (LoginRequest) TableName() string {
	return "staff"
}

func (r *LoginRequest) Validate() (err error) {
	validate := validator.New()
	if err = validate.Struct(r); err != nil {
		return
	}
	users := make([]LoginRequest, 0)
	err = models.Model.Where("name = ?", r.Username).Find(&users).Error
	for _, user := range users {
		if user.Password == r.Password {
			*r = user
			return
		}
	}
	return errno.ErrUserAuthFailed
}

func Login(c *gin.Context) {
	authDr, _ := c.MustGet("jwt_auth").(*auth.Auth)
	var r LoginRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Validate the data.
	if err := r.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrUserAuthFailed, nil)
		return
	}

	staff, _ := models.RetrieveStaff(r.Id)
	staffMap := map[string]interface{}{
		"id":            staff.Id,
		"name":          staff.Name,
		"permission":    staff.Permission,
		"specialty_id":  staff.SpecialtyId,
		"department_id": staff.DepartmentId,
	}
	token, _ := (*authDr).Login(c.Request, c.Writer, staffMap).(string)

	handler.SendResponse(c, nil, map[string]interface{}{"token": token})
}
