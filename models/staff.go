package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	"pppobear.cn/jxc-backend/module/handler"
)

var (
	StaffSearchFields = []string{
		"id",
		"name",
		"specialty_id--id:specialty.id&name",
		"department_id--id:department.id&name&head_name",
	}
	SpeSearchFields = []string{
		"id",
		"name",
	}
	DepSearchFields = []string{
		"id",
		"name",
		"head_name",
	}
)

type StaffFilterRequest struct {
	Gender       string       `form:"gender"`
	SpecialtyId  uint64       `form:"specialty_id"`
	DepartmentId uint64       `form:"department_id"`
	Married      sql.NullBool `form:"married"`
	Permission   string       `form:"permission"`
	StartDate    string       `form:"start_date"`
	EndDate      string       `form:"end_date"`
}

func (sfr *StaffFilterRequest) Filter(sql *gorm.DB) {
	if sfr.Gender != "" {
		*sql = *sql.Where("gender = ?", sfr.Gender)
	}
	if sfr.SpecialtyId != 0 {
		*sql = *sql.Where("specialty_id = ?", sfr.SpecialtyId)
	}
	if sfr.DepartmentId != 0 {
		*sql = *sql.Where("department_id = ?", sfr.DepartmentId)
	}
	if sfr.Married.Valid {
		*sql = *sql.Where("married = ?", sfr.Married.Bool)
	}
	if sfr.Permission != "" {
		*sql = *sql.Where("permission = ?", sfr.Permission)
	}

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
			*sql = *sql.Where("birthday <= ?", EndDate)
		}
	} else {
		if EndDate == nil {
			*sql = *sql.Where("birthday >= ?", StartDate)
		} else {
			*sql = *sql.Where("birthday BETWEEN ? AND ?", StartDate, EndDate)
		}
	}
}

type SpecialtyModel struct {
	Id   uint64 `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

type DepartmentModel struct {
	Id       uint64 `gorm:"primary_key" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	HeadName string `gorm:"not null" json:"head_name"`
}

type StaffModel struct {
	Id           string           `gorm:"primary_key" json:"id"`
	Name         string           `gorm:"not null" json:"name"`
	Gender       string           `gorm:"size:2" json:"gender"`
	Birthday     handler.JsonDate `json:"birthday"`
	Specialty    SpecialtyModel   `json:"specialty"`
	SpecialtyId  uint64           `json:"-"`
	Department   DepartmentModel  `json:"department"`
	DepartmentId uint64           `json:"-"`
	Married      sql.NullBool     `json:"married"`
	Permission   uint             `json:"permission"`
}

func (StaffModel) TableName() string {
	return "staff"
}

func (SpecialtyModel) TableName() string {
	return "specialty"
}

func (DepartmentModel) TableName() string {
	return "department"
}

func (s *StaffModel) Create() error {
	return Model.Create(s).Set("gorm:auto_preload", true).First(s).Error
}

func ListStaff(req *handler.ListRequest, filter Filter) (handler.ListResponse, error) {
	staffs := make([]*StaffModel, 0)
	return List(staffs, new(StaffModel), req, filter, nil)
}

func RetrieveStaff(id string) (*StaffModel, error) {
	staff := &StaffModel{}
	r := Model.Where("id = ?", id).Set("gorm:auto_preload", true).First(staff)
	return staff, r.Error
}

func (s *StaffModel) Update() (err error) {
	err = Model.Model(s).
		Select("name", "gender", "birthday", "married").
		Updates(s).Error
	if err != nil {
		return
	}
	return Model.Where("id = ?", s.Id).First(s).Error
}

func DeleteStaff(id string) error {
	return Model.Delete(&StaffModel{Id: id}).Error
}

func (spe *SpecialtyModel) Create() error {
	return Model.Create(spe).First(spe).Error
}

func ListSpecialty(req *handler.ListRequest) (handler.ListResponse, error) {
	specialties := make([]*SpecialtyModel, 0)
	return List(specialties, new(SpecialtyModel), req, nil, nil)
}

func RetrieveSpecialty(id string) (*SpecialtyModel, error) {
	specialty := &SpecialtyModel{}
	r := Model.Where("id = ?", id).First(specialty)
	return specialty, r.Error
}

func (spe *SpecialtyModel) Update() (err error) {
	err = Model.Model(spe).Select("name").Updates(spe).Error
	if err != nil {
		return
	}
	return Model.Where("id = ?", spe.Id).First(spe).Error
}

func DeleteSpecialty(idStr string) error {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return err
	}
	return Model.Delete(&SpecialtyModel{Id: id}).Error
}

func (dep *DepartmentModel) Create() error {
	return Model.Create(dep).First(dep).Error
}

func ListDepartment(req *handler.ListRequest) (handler.ListResponse, error) {
	departments := make([]*DepartmentModel, 0)
	return List(departments, new(DepartmentModel), req, nil, nil)
}

func RetrieveDepartment(id string) (*DepartmentModel, error) {
	department := &DepartmentModel{}
	r := Model.Where("id = ?", id).First(department)
	return department, r.Error
}

func (dep *DepartmentModel) Update() (err error) {
	err = Model.Model(dep).Select("name", "head_name").Updates(dep).Error
	if err != nil {
		return
	}
	return Model.Where("id = ?", dep.Id).First(dep).Error
}

func DeleteDepartment(idStr string) error {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return err
	}
	return Model.Delete(&DepartmentModel{Id: id}).Error
}
