package models

import "pppobear.cn/jxc-backend/module/handler"

var (
	CustomerSearchFields = []string{
		"id",
		"name",
		"contact",
		"phone",
	}
)

type CustomerModel struct {
	Id      string `gorm:"primary_key;column:id" json:"id" binding:"required" validate:"min=1,max=10"`
	Name    string `gorm:"not null" json:"name" binding:"required" validate:"min=1,max=20"`
	Contact string `gorm:"not null" json:"contact" binding:"required" validate:"min=1,max=8"`
	Phone   string `gorm:"not null" json:"phone" binding:"required" validate:"min=1,max=20"`
}

func (CustomerModel) TableName() string {
	return "customers"
}

func ListCustomer(req *handler.ListRequest) (handler.ListResponse, error) {
	goods := make([]*CustomerModel, 0)
	return List(goods, &CustomerModel{}, req, nil, nil)
}

func RetrieveCustomer(id string) (*CustomerModel, error) {
	c := &CustomerModel{}
	r := Model.Where("id = ?", id).First(&c)
	return c, r.Error
}

func (c *CustomerModel) Create() error {
	return Model.Create(c).First(c).Error
}

func DeleteCustomer(id string) error {
	return Model.Delete(&CustomerModel{Id: id}).Error
}

func (c *CustomerModel) Update() (err error) {
	err = Model.Model(c).Select("name", "contact", "phone").Updates(c).Error
	if err != nil {
		return
	}
	return Model.Where("id = ?", c.Id).First(c).Error
}
