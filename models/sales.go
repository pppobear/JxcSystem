package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	"pppobear.cn/jxc-backend/module/handler"
)

var (
	SalesSearchFields = []string{
		"id",
		"customer_id--id:customers.id&name",
		"sales_staff_id--id:staff.id&name",
	}
)

type SalesFilterRequest struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}

func (sfr *SalesFilterRequest) Filter(sql *gorm.DB) {
	var (
		StartDate *time.Time
		EndDate   *time.Time
	)
	sd, err := time.Parse(`2006-01-02`, sfr.StartDate)
	if err != nil {
		StartDate = nil
	} else {
		StartDate = &sd
	}
	ed, err := time.Parse(`2006-01-02`, sfr.EndDate)
	if err != nil {
		EndDate = nil
	} else {
		ed = time.Time.Add(ed, time.Hour*24-time.Nanosecond)
		EndDate = &ed
	}

	if StartDate == nil {
		if EndDate != nil {
			*sql = *sql.Where("datetime <= ?", EndDate)
		}
	} else {
		if EndDate == nil {
			*sql = *sql.Where("datetime >= ?", StartDate)
		} else {
			*sql = *sql.Where("datetime BETWEEN ? AND ?", StartDate, EndDate)
		}
	}
}

type SalesModel struct {
	Id           uint64             `gorm:"primary_key" json:"id"`
	Datetime     handler.JsonDate   `gorm:"not null" json:"datetime"`
	Customer     CustomerModel      `json:"customer"`
	CustomerId   string             `json:"-"`
	SalesStaff   StaffModel         `json:"sales_staff"`
	SalesStaffId string             `json:"-"`
	Details      []SalesDetailModel `gorm:"foreignkey:SalesId;association_foreignkey:Id" json:"details"`
}

type SalesDetailModel struct {
	SalesId   uint64     `gorm:"primary_key" json:"-"`
	GoodsId   string     `gorm:"primary_key" json:"-"`
	Goods     GoodsModel `json:"goods"`
	Number    uint       `json:"number"`
	UnitPrice float64    `json:"unit_price"`
}

func (SalesModel) TableName() string {
	return "sales"
}

func (sales *SalesModel) Create() error {
	return Model.Create(sales).Set("gorm:auto_preload", true).First(sales).Error
}

func (SalesDetailModel) TableName() string {
	return "sales_detail"
}

func (salesD *SalesDetailModel) BeforeCreate() (err error) {
	curInv, _ := RetrieveInventory(salesD.GoodsId)
	salesD.UnitPrice = curInv.UnitPrice
	return
}

func (salesD *SalesDetailModel) Create() error {
	return Model.Create(salesD).Set("gorm:auto_preload", true).First(salesD).Error
}

func (sales *SalesModel) Update() (err error) {
	err = Model.Model(sales).Select("customer_id", "sales_staff_id", "datetime").Updates(sales).Error
	if err != nil {
		return
	}
	return Model.Where("id = ?", sales.Id).Set("gorm:auto_preload", true).First(sales).Error
}

func (salesD *SalesDetailModel) Update() (err error) {
	err = Model.Model(salesD).Select("number").Updates(salesD).Error
	if err != nil {
		return
	}
	return Model.Where("sales_id = ? and goods_id = ?", salesD.SalesId, salesD.GoodsId).
		Set("gorm:auto_preload", true).First(salesD).Error
}

func DeleteSales(idStr string) error {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return err
	}
	return Model.Delete(&SalesModel{Id: id}).Error
}

func DeleteSalesDetail(id uint64, gid string) error {
	return Model.Delete(&SalesDetailModel{SalesId: id, GoodsId: gid}).Error
}

func ListSales(req *handler.ListRequest, filter Filter, user *map[string]interface{}) (handler.ListResponse, error) {
	sales := make([]*SalesModel, 0)
	return List(sales, new(SalesModel), req, filter, user)
}

func RetrieveSales(id string) (*SalesModel, error) {
	s := &SalesModel{}
	r := Model.Where("id = ?", id).Set("gorm:auto_preload", true).First(&s)
	return s, r.Error
}
