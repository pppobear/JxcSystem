package models

import (
	"regexp"

	"github.com/jinzhu/gorm"

	"pppobear.cn/jxc-backend/module/handler"
)

var (
	GoodsSearchFields = []string{
		"id",
		"name",
		"model",
		"specifications",
	}
	InventorySearchFields = []string{
		"unit_price",
		"id--id:goods.id&name&model&specifications",
	}
)

type GoodsFilterRequest struct {
	Specifications string `form:"specifications"`
	UnitName       string `form:"unit_name"`
}

func (gfr GoodsFilterRequest) Filter(sql *gorm.DB) {
	if gfr.Specifications != "" {
		*sql = *sql.Where("specifications LIKE ?", "%"+gfr.Specifications+"%")
	}
	if gfr.UnitName != "" {
		*sql = *sql.Where("unit_name LIKE ?", "%"+gfr.UnitName+"%")
	}
}

type InventoryFilterRequest struct {
	MinNumber    string `form:"min_number"`
	MaxNumber    string `form:"max_number"`
	MinUnitPrice string `form:"min_unit_price"`
	MaxUnitPrice string `form:"max_unit_price"`
}

func (ifr *InventoryFilterRequest) Filter(sql *gorm.DB) {
	if regexp.MustCompile(`\D`).MatchString(ifr.MinNumber) {
		ifr.MinNumber = ""
	}
	if regexp.MustCompile(`\D`).MatchString(ifr.MaxNumber) {
		ifr.MaxNumber = ""
	}
	if regexp.MustCompile(`\D`).MatchString(ifr.MinUnitPrice) {
		ifr.MinUnitPrice = ""
	}
	if regexp.MustCompile(`\D`).MatchString(ifr.MaxUnitPrice) {
		ifr.MaxUnitPrice = ""
	}

	if ifr.MinNumber == "" {
		if ifr.MaxNumber != "" {
			*sql = *sql.Where("number <= ?", ifr.MaxNumber)
		}
	} else {
		if ifr.MaxNumber == "" {
			*sql = *sql.Where("number >= ?", ifr.MinNumber)
		} else {
			*sql = *sql.Where("number BETWEEN ? AND ?", ifr.MinNumber, ifr.MaxNumber)
		}
	}
	if ifr.MinUnitPrice == "" {
		if ifr.MaxUnitPrice != "" {
			*sql = *sql.Where("unit_price <= ?", ifr.MaxUnitPrice)
		}
	} else {
		if ifr.MaxUnitPrice == "" {
			*sql = *sql.Where("unit_price >= ?", ifr.MinUnitPrice)
		} else {
			*sql = *sql.Where("unit_price BETWEEN ? AND ?", ifr.MinUnitPrice, ifr.MaxUnitPrice)
		}
	}
}

type GoodsModel struct {
	Id             string `gorm:"primary_key" json:"id"`
	Name           string `gorm:"not null" json:"name"`
	Model          string `gorm:"not null" json:"model"`
	Specifications string `json:"specifications"`
	UnitName       string `gorm:"not null" json:"unit_name"`
	MaxInventory   uint   `json:"max_inventory"`
	MinInventory   uint   `json:"min_inventory"`
}

type InventoryModel struct {
	Id        string     `gorm:"primary_key" json:"-"`
	Goods     GoodsModel `gorm:"foreignkey:Id";association_foreignkey:Id" json:"goods"`
	Number    uint       `json:"number"`
	UnitPrice float64    `json:"unit_price"`
}

func (GoodsModel) TableName() string {
	return "goods"
}

func (InventoryModel) TableName() string {
	return "inventory"
}

func ListGoods(req *handler.ListRequest, filter Filter) (handler.ListResponse, error) {
	goods := make([]*GoodsModel, 0)
	return List(goods, &GoodsModel{}, req, filter, nil)
}

func ListInventory(req *handler.ListRequest, filter Filter) (handler.ListResponse, error) {
	inventories := make([]*InventoryModel, 0)
	return List(inventories, &InventoryModel{}, req, filter, nil)
}

func RetrieveGoods(id string) (*GoodsModel, error) {
	g := &GoodsModel{}
	r := Model.Where("id = ?", id).First(&g)
	return g, r.Error
}

func RetrieveInventory(id string) (*InventoryModel, error) {
	i := &InventoryModel{}
	r := Model.Set("gorm:auto_preload", true).Where("id = ?", id).First(&i)
	return i, r.Error
}

func (g *GoodsModel) Create() error {
	sql := Model.Model(g)
	var omits []string
	if g.MaxInventory == 0 {
		omits = append(omits, "min_inventory", "max_inventory")
	}
	if len(g.Specifications) == 0 {
		omits = append(omits, "specifications")
	}
	sql = sql.Omit(omits...)
	return sql.Create(g).First(g).Error
}

func (i *InventoryModel) Create() error {
	sql := Model.Model(i)
	if i.UnitPrice == 0 {
		sql = sql.Omit("unit_price")
	}
	return sql.Create(i).First(i).Error
}

func DeleteGoods(id string) error {
	return Model.Delete(&GoodsModel{Id: id}).Error
}

func (g *GoodsModel) Update() (err error) {
	err = Model.Model(g).
		Select("name", "model", "specifications", "unit_name", "max_inventory", "min_inventory").
		Updates(g).Error
	if err != nil {
		return
	}
	return Model.Where("id = ?", g.Id).First(g).Error
}

func (i *InventoryModel) Update() (err error) {
	err = Model.Model(i).
		Select("number", "unit_price").
		Updates(i).Error
	if err != nil {
		return
	}
	return Model.Where("id = ?", i.Id).Set("gorm:auto_preload", true).First(i).Error
}
