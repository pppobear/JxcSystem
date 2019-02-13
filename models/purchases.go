package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	"pppobear.cn/jxc-backend/module/handler"
)

var (
	PurchasesSearchFields = []string{
		"id",
		"supplier_id--id:customers.id&name",
		"buyer_id--id:staff.id&name",
	}
)

type PurchasesFilterRequest struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}

func (pfr *PurchasesFilterRequest) Filter(sql *gorm.DB) {
	var (
		StartDate *time.Time
		EndDate   *time.Time
	)
	sd, err := time.Parse(`2006-01-02`, pfr.StartDate)
	if err != nil {
		StartDate = nil
	} else {
		StartDate = &sd
	}
	ed, err := time.Parse(`2006-01-02`, pfr.EndDate)
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

type PurchasesModel struct {
	Id         uint64                 `gorm:"primary_key" json:"id"`
	Datetime   handler.JsonDate       `gorm:"not null" json:"datetime"`
	Supplier   CustomerModel          `json:"supplier"`
	SupplierId string                 `json:"-"`
	Buyer      StaffModel             `json:"buyer"`
	BuyerId    string                 `json:"-"`
	Details    []PurchasesDetailModel `gorm:"foreignkey:PurchaseID;association_foreignkey:Id" json:"details"`
}

type PurchasesDetailModel struct {
	PurchaseID uint64     `gorm:"primary_key" json:"-"`
	GoodsId    string     `gorm:"primary_key" json:"-"`
	Goods      GoodsModel `json:"goods"`
	Number     uint       `json:"number"`
	UnitPrice  float64    `json:"unit_price"`
}

func (PurchasesModel) TableName() string {
	return "purchases"
}

func (PurchasesDetailModel) TableName() string {
	return "purchases_detail"
}

func (pur *PurchasesModel) Create() error {
	return Model.Create(pur).Set("gorm:auto_preload", true).First(pur).Error
}

//func (purD *PurchasesDetailModel) BeforeCreate() (err error) {
//	curInv, _ := RetrieveInventory(purD.GoodsId)
//	purD.UnitPrice = curInv.UnitPrice
//	return
//}

func (purD *PurchasesDetailModel) Create() error {
	return Model.Create(purD).Set("gorm:auto_preload", true).First(purD).Error
}

func (pur *PurchasesModel) Update() (err error) {
	err = Model.Model(pur).Select("supplier_id", "buyer_id", "datetime").Updates(pur).Error
	if err != nil {
		return
	}
	return Model.Where("id = ?", pur.Id).Set("gorm:auto_preload", true).First(pur).Error
}

func (purD *PurchasesDetailModel) Update() (err error) {
	err = Model.Model(purD).Select("number").Updates(purD).Error
	if err != nil {
		return
	}
	return Model.Where("purchase_id = ? and goods_id = ?", purD.PurchaseID, purD.GoodsId).
		Set("gorm:auto_preload", true).First(purD).Error
}

func DeletePurchases(idStr string) error {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return err
	}
	return Model.Delete(&PurchasesModel{Id: id}).Error
}

func DeletePurchasesDetail(id uint64, gid string) error {
	return Model.Delete(&PurchasesDetailModel{PurchaseID: id, GoodsId: gid}).Error
}

func ListPurchases(req *handler.ListRequest, filter Filter, user *map[string]interface{}) (handler.ListResponse, error) {
	purchases := make([]*PurchasesModel, 0)
	return List(purchases, &PurchasesModel{}, req, filter, user)
}

func RetrievePurchase(id string) (*PurchasesModel, error) {
	p := &PurchasesModel{}
	r := Model.Where("id = ?", id).Set("gorm:auto_preload", true).First(&p)
	return p, r.Error
}
