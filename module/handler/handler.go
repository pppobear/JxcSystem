package handler

import (
	"database/sql/driver"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/lexkong/log"
	"gopkg.in/go-playground/validator.v9"

	"pppobear.cn/jxc-backend/module/errno"
)

type JsonDate struct {
	time.Time
}

func (d *JsonDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(d.Time).Format("2006-01-02"))), nil
}

func (d *JsonDate) UnmarshalJSON(data []byte) error {
	var err error
	d.Time, err = time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return err
	}
	return nil
}

func (d JsonDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if d.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return d.Time, nil
}

func (d *JsonDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*d = JsonDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (d JsonDate) String() string {
	return time.Time(d.Time).Format("2006-01-02")
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListRequest struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	OrderBy  string `form:"order_by"`
	Search   string `form:"search"`
}

type ListResponse struct {
	Page      uint64      `json:"page"`
	PageSize  uint64      `json:"page_size"`
	Total     uint64      `json:"total"`
	TotalPage uint64      `json:"total_page"`
	PrePage   string      `json:"pre_page"`
	NextPage  string      `json:"next_page"`
	Data      interface{} `json:"data"`
}

type AddDetailRequest struct {
	Id        uint64  `json:"-" validate:"min=1"`
	UnitPrice float64 `json:"unit_price" binding:"required" validate:"min=1"`
	GoodsId   string  `json:"goods_id" binding:"required" validate:"min=1"`
	Number    uint    `json:"number" binding:"required" validate:"min=1"`
}

type CreatePurSalRequest struct {
	Datetime   JsonDate               `json:"datetime" binding:"required"`
	CustomerId string                 `json:"customer_id" binding:"required" validate:"min=1"`
	StaffId    string                 `json:"staff_id" binding:"required" validate:"min=1"`
	Details    []*CreateDetailRequest `json:"details"`
}

type CreateDetailRequest struct {
	GoodsId   string  `json:"goods_id" binding:"required" validate:"min=1"`
	UnitPrice float64 `json:"unit_price" binding:"required" validate:"min=1"`
	Number    uint    `json:"number" binding:"required" validate:"min=1"`
}

type UpdateDetailRequest struct {
	Number uint `json:"number" binding:"required" validate:"min=1"`
}

type ValidateRequest interface {
	Validate() error
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func ValidateAdapt(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}

func (r *CreatePurSalRequest) Validate() error {
	return ValidateAdapt(r)
}

func (r *AddDetailRequest) Validate() error {
	return ValidateAdapt(r)
}

func (r *CreateDetailRequest) Validate() error {
	return ValidateAdapt(r)
}

func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				log.Error("", err.(error))

				var (
					errMsg     string
					mysqlError *mysql.MySQLError
					ok         bool
				)
				if errMsg, ok = err.(string); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg":  "system error, " + errMsg,
					})
					return
				} else if mysqlError, ok = err.(*mysql.MySQLError); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg":  "system error, " + mysqlError.Error(),
					})
					return
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 500,
						"msg":  "system error",
					})
					return
				}
			}
		}()
		c.Next()
	}
}
